package state

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwftest"
	"github.com/stretchr/testify/assert"
)

func TestSignalOrderPPKState_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	emptyObj := iwftest.NewTestObject(nil)

	state := NewSignalOrderPPKState()

	cmdReq, err := state.Start(mockWorkflowCtx, emptyObj, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.AllCommandsCompletedRequest(
		iwf.NewSignalCommand("", PPK_CHANNEL),
	), cmdReq)
}

func TestSignalOrderPPKStateToComplete_Decide(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	emptyObj := iwftest.NewTestObject(nil)
	cmdResults := iwf.CommandResults{
		Signals: []iwf.SignalCommandResult{
			{
				ChannelName: PPK_CHANNEL,
				SignalValue: iwftest.NewTestObject("completed"),
				Status:      iwfidl.RECEIVED,
			},
		},
	}

	state := NewSignalOrderPPKState()

	decision, err := state.Decide(mockWorkflowCtx, emptyObj, cmdResults, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.SingleNextState(&setOrderToCompleted{}, nil), decision)
}

func TestSignalOrderPPKStateToRejected_Decide(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	emptyObj := iwftest.NewTestObject(nil)
	cmdResults := iwf.CommandResults{
		Signals: []iwf.SignalCommandResult{
			{
				ChannelName: PPK_CHANNEL,
				SignalValue: iwftest.NewTestObject("rejected"),
				Status:      iwfidl.RECEIVED,
			},
		},
	}

	state := NewSignalOrderPPKState()

	decision, err := state.Decide(mockWorkflowCtx, emptyObj, cmdResults, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.SingleNextState(&setOrderToRejected{}, nil), decision)
}
