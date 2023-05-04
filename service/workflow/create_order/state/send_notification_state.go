package state

import (
	"fmt"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type sendNotificationState struct {
	iwf.DefaultStateIdAndOptions

	repo repository.OrderRepository
}

func NewSendNotificationState(repo repository.OrderRepository) iwf.WorkflowState {
	return &sendNotificationState{
		repo: repo,
	}
}

func (ths *sendNotificationState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	// Send notification via email or else
	var orderId string
	persistence.GetDataObject(KeyOrderId, &orderId)

	order, err := ths.repo.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	fmt.Println(order)
	return iwf.EmptyCommandRequest(), nil
}

func (ths *sendNotificationState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.MultiNextStates(&expiredTimerState{}, &signalOrderPPKState{}), nil
}
