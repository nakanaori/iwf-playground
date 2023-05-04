package service

import (
	"context"
	"iwf-playground/model"
	"iwf-playground/repository"
	"iwf-playground/service/workflow/create_order"
	"iwf-playground/service/workflow/create_order/state"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, orderId string) (model.Order, error)
	CheckoutOrder(ctx context.Context, orderId string) (string, error)
	RejectOrder(ctx context.Context, orderId string) error
	AcceptOrder(ctx context.Context, orderId string) error
	OrderInvokeStartHandler(ctx context.Context, req iwfidl.WorkflowStateStartRequest) (*iwfidl.WorkflowStateStartResponse, error)
	OrderInvokeDecideHandler(ctx context.Context, req iwfidl.WorkflowStateDecideRequest) (*iwfidl.WorkflowStateDecideResponse, error)
}

type service struct {
	OrderRepository repository.OrderRepository
	IwfClient       iwf.Client
	WorkerService   iwf.WorkerService
}

func NewOrderService() OrderService {
	repo := repository.NewOrderRepository()
	registry, _ := create_order.GetRegistry()
	return &service{
		OrderRepository: repo,
		IwfClient:       iwf.NewClient(registry, nil),
		WorkerService:   iwf.NewWorkerService(registry, nil),
	}
}

func (ths *service) generateWfId(orderId string) string {
	return "CREATE_ORDER_" + orderId
}

func (ths *service) CheckoutOrder(ctx context.Context, orderId string) (string, error) {
	createOrderWorkflow := create_order.NewCreateOrderWorkflow(ths.OrderRepository)
	wfId := ths.generateWfId(orderId)
	// 3600 = workflow timeout in seconds
	// orderId = input for workflow.
	runId, err := ths.IwfClient.StartWorkflow(ctx, createOrderWorkflow, wfId, 3600, orderId, nil)
	if err != nil {
		return "", err
	}

	return runId, nil
}

func (ths *service) RejectOrder(ctx context.Context, orderId string) error {
	wfid := ths.generateWfId(orderId)
	err := ths.IwfClient.SignalWorkflow(ctx, create_order.NewCreateOrderWorkflow(ths.OrderRepository), wfid, "", state.PPK_CHANNEL, model.CANCELLED)
	if err != nil {
		return err
	}
	return nil
}

func (ths *service) AcceptOrder(ctx context.Context, orderId string) error {
	wfid := ths.generateWfId(orderId)
	err := ths.IwfClient.SignalWorkflow(ctx, create_order.NewCreateOrderWorkflow(ths.OrderRepository), wfid, "", state.PPK_CHANNEL, model.COMPLETED)
	if err != nil {
		return err
	}
	return nil
}

func (ths *service) OrderInvokeStartHandler(ctx context.Context, req iwfidl.WorkflowStateStartRequest) (*iwfidl.WorkflowStateStartResponse, error) {
	return ths.WorkerService.HandleWorkflowStateStart(ctx, req)
}

func (ths *service) OrderInvokeDecideHandler(ctx context.Context, req iwfidl.WorkflowStateDecideRequest) (*iwfidl.WorkflowStateDecideResponse, error) {
	return ths.WorkerService.HandleWorkflowStateDecide(ctx, req)
}

func (ths *service) GetOrderByID(ctx context.Context, orderId string) (model.Order, error) {
	return ths.OrderRepository.GetOrderById(orderId)
}
