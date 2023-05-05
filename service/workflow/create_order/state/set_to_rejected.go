package state

import (
	"iwf-playground/model"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type setOrderToRejected struct {
	iwf.DefaultStateIdAndOptions

	repo repository.OrderRepository
}

func NewSetOrderToRejected(repo repository.OrderRepository) iwf.WorkflowState {
	return &setOrderToRejected{
		repo: repo,
	}
}

func (ths *setOrderToRejected) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var orderId string
	persistence.GetDataObject(KeyOrderId, &orderId)

	// Set Order to Rejected
	err := ths.repo.SetOrderStatus(orderId, string(model.REJECTED))
	if err != nil {
		return nil, err
	}

	return iwf.EmptyCommandRequest(), nil
}

func (ths *setOrderToRejected) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceCompleteWorkflow(nil), nil
}
