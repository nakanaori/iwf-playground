package state

import (
	"iwf-playground/model"
	"iwf-playground/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwftest"
	"github.com/stretchr/testify/assert"
)

func TestSendNotificationState_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockOrderRepository(ctrl)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	orderId := "test-order-1"
	emptyObj := iwftest.NewTestObject(nil)

	state := NewSendNotificationState(mockRepo)

	mockPersistence.EXPECT().GetDataObject(KeyOrderId, gomock.Any()).SetArg(1, orderId)
	mockRepo.EXPECT().GetOrderById(orderId).Return(model.Order{}, nil)

	cmdReq, err := state.Start(mockWorkflowCtx, emptyObj, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.EmptyCommandRequest(), cmdReq)
}

func TestSendNotificationState_Decide(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockOrderRepository(ctrl)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	testOrderIdObj := iwftest.NewTestObject(nil)
	emptyCmdResults := iwf.CommandResults{}

	state := NewSendNotificationState(mockRepo)

	decision, err := state.Decide(mockWorkflowCtx, testOrderIdObj, emptyCmdResults, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.MultiNextStates(&expiredTimerState{}, &signalOrderPPKState{}), decision)
}
