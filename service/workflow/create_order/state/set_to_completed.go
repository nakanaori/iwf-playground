package state

import (
	"iwf-playground/model"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type setOrderToCompleted struct {
	iwf.DefaultStateIdAndOptions

	repo repository.OrderRepository
}

func NewSetOrderToCompleted(repo repository.OrderRepository) iwf.WorkflowState {
	return &setOrderToCompleted{
		repo: repo,
	}
}

func (ths *setOrderToCompleted) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var orderId string
	persistence.GetDataObject(KeyOrderId, &orderId)

	// Set Order to Completed
	err := ths.repo.SetOrderStatus(orderId, string(model.COMPLETED))
	if err != nil {
		return nil, err
	}

	return iwf.EmptyCommandRequest(), nil
}

func (ths *setOrderToCompleted) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceCompleteWorkflow(nil), nil
}
