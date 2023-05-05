package state

import (
	"iwf-playground/model"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type setOrderToCancelled struct {
	iwf.DefaultStateIdAndOptions

	repo repository.OrderRepository
}

func NewSetOrderToCancelled(repo repository.OrderRepository) iwf.WorkflowState {
	return &setOrderToCancelled{
		repo: repo,
	}
}

func (ths *setOrderToCancelled) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var orderId string
	persistence.GetDataObject(KeyOrderId, &orderId)

	// Set Order to Cancelled
	err := ths.repo.SetOrderStatus(orderId, string(model.CANCELLED))
	if err != nil {
		return nil, err
	}

	return iwf.EmptyCommandRequest(), nil
}

func (ths *setOrderToCancelled) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceCompleteWorkflow(nil), nil
}
