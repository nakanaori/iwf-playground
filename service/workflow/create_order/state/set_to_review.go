package state

import (
	"iwf-playground/model"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type setOrderToReview struct {
	iwf.DefaultStateIdAndOptions

	repo repository.OrderRepository
}

func NewSetOrderToReview(repo repository.OrderRepository) iwf.WorkflowState {
	return &setOrderToReview{
		repo: repo,
	}
}

func (ths *setOrderToReview) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var orderId string
	persistence.GetDataObject(KeyOrderId, &orderId)

	// Set Order to expired
	err := ths.repo.SetOrderStatus(orderId, string(model.ON_REVIEW))
	if err != nil {
		return nil, err
	}

	return iwf.EmptyCommandRequest(), nil
}

func (ths *setOrderToReview) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.SingleNextState(&sendNotificationState{}, nil), nil
}
