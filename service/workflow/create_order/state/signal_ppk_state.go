package state

import (
	"fmt"
	"iwf-playground/model"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type signalOrderPPKState struct {
	iwf.DefaultStateIdAndOptions
}

func NewSignalOrderPPKState() iwf.WorkflowState {
	return &signalOrderPPKState{}
}

func (ths *signalOrderPPKState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.AnyCommandCompletedRequest(
		iwf.NewSignalCommand("", PPK_CHANNEL),
	), nil
}

func (b signalOrderPPKState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	signal := commandResults.Signals[0]
	if signal.Status == iwfidl.RECEIVED {
		var status string
		signal.SignalValue.Get(&status)
		switch status {
		case string(model.REJECTED):
			return iwf.SingleNextState(&setOrderToRejected{}, nil), nil
		case string(model.COMPLETED):
			return iwf.SingleNextState(&setOrderToCompleted{}, nil), nil
		}
	}

	return nil, fmt.Errorf(PPK_CHANNEL + " doesn't receive correct value")
}
