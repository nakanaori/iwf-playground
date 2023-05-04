package state

import (
	"iwf-playground/model"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type setOrderToExpired struct {
	iwf.DefaultStateIdAndOptions

	repo repository.OrderRepository
}

func NewSetOrderToExpired(repo repository.OrderRepository) iwf.WorkflowState {
	return &setOrderToExpired{
		repo: repo,
	}
}

func (ths *setOrderToExpired) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var orderId string
	persistence.GetDataObject(KeyOrderId, &orderId)

	// Set Order to expired
	err := ths.repo.SetOrderStatus(orderId, string(model.CANCELLED))
	if err != nil {
		return nil, err
	}

	return iwf.EmptyCommandRequest(), nil
}

func (ths *setOrderToExpired) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceCompleteWorkflow(nil), nil
}
