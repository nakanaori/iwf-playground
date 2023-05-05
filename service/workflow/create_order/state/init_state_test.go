package state

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwftest"
	"github.com/stretchr/testify/assert"
)

func TestInitState_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	orderId := "test-order-1"
	testOrderIdObj := iwftest.NewTestObject(orderId)

	state := NewInitState()

	mockPersistence.EXPECT().SetDataObject(KeyOrderId, orderId)

	cmdReq, err := state.Start(mockWorkflowCtx, testOrderIdObj, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.EmptyCommandRequest(), cmdReq)
}

func TestInitState_Decide(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	orderId := "test-order-1"
	testOrderIdObj := iwftest.NewTestObject(orderId)
	emptyCmdResults := iwf.CommandResults{}

	state := NewInitState()

	decision, err := state.Decide(mockWorkflowCtx, testOrderIdObj, emptyCmdResults, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.SingleNextState(&setOrderToReview{}, nil), decision)
}
