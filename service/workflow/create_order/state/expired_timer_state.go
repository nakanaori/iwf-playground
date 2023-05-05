package state

import (
	"time"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type expiredTimerState struct {
	iwf.DefaultStateIdAndOptions
}

func NewExpiredTimerState() iwf.WorkflowState {
	return &expiredTimerState{}
}

func (ths *expiredTimerState) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.AllCommandsCompletedRequest(
		iwf.NewTimerCommand("", time.Now().Add(ExpiredTime)),
	), nil
}

func (ths *expiredTimerState) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.SingleNextState(&setOrderToCancelled{}, nil), nil
}
