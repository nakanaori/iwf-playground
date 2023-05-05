package state

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwftest"
	"github.com/stretchr/testify/assert"
)

func TestExpiredTimerState_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	emptyObj := iwftest.NewTestObject(nil)

	state := NewExpiredTimerState()

	cmdReq, err := state.Start(mockWorkflowCtx, emptyObj, mockPersistence, mockCommunication)

	firingTime := cmdReq.Commands[0].TimerCommand.FiringUnixTimestampSeconds

	assert.Nil(t, err)
	assert.Equal(t, iwf.AllCommandsCompletedRequest(
		iwf.NewTimerCommand("", time.Unix(firingTime, 0)),
	), cmdReq)
}

func TestExpiredTimerStateToComplete_Decide(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	emptyObj := iwftest.NewTestObject(nil)
	cmdResults := iwf.CommandResults{}

	state := NewExpiredTimerState()

	decision, err := state.Decide(mockWorkflowCtx, emptyObj, cmdResults, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.SingleNextState(&setOrderToCancelled{}, nil), decision)
}
