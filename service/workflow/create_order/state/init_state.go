package state

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type initState struct {
	iwf.DefaultStateIdAndOptions
}

func NewInitState() iwf.WorkflowState {
	return &initState{}
}

func (ths *initState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var orderId string
	input.Get(&orderId)

	persistence.SetDataObject(KeyOrderId, orderId)

	return iwf.EmptyCommandRequest(), nil
}

func (ths *initState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.SingleNextState(&setOrderToReview{}, nil), nil
}
