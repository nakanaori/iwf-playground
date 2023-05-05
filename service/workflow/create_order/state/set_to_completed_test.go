package state

import (
	"fmt"
	"iwf-playground/model"
	"iwf-playground/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwftest"
	"github.com/stretchr/testify/assert"
)

func TestSetToCompleted_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockOrderRepository(ctrl)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	orderId := "test-order-1"
	emptyObj := iwftest.NewTestObject(nil)

	state := NewSetOrderToCompleted(mockRepo)

	mockPersistence.EXPECT().GetDataObject(KeyOrderId, gomock.Any()).SetArg(1, orderId)
	mockRepo.EXPECT().SetOrderStatus(orderId, string(model.COMPLETED)).Return(nil)

	cmdReq, err := state.Start(mockWorkflowCtx, emptyObj, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.EmptyCommandRequest(), cmdReq)

	mockPersistence.EXPECT().GetDataObject(KeyOrderId, gomock.Any()).SetArg(1, orderId)
	mockRepo.EXPECT().SetOrderStatus(orderId, string(model.COMPLETED)).Return(fmt.Errorf("error"))

	_, err = state.Start(mockWorkflowCtx, emptyObj, mockPersistence, mockCommunication)

	assert.NotNil(t, err)
}

func TestSetToCompleted_Decide(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockOrderRepository(ctrl)
	mockPersistence := iwftest.NewMockPersistence(ctrl)
	mockWorkflowCtx := iwftest.NewMockWorkflowContext(ctrl)
	mockCommunication := iwftest.NewMockCommunication(ctrl)

	emptyObj := iwftest.NewTestObject(nil)
	emptyCmdResults := iwf.CommandResults{}

	state := NewSetOrderToCompleted(mockRepo)

	decision, err := state.Decide(mockWorkflowCtx, emptyObj, emptyCmdResults, mockPersistence, mockCommunication)

	assert.Nil(t, err)
	assert.Equal(t, iwf.ForceCompleteWorkflow(nil), decision)
}
