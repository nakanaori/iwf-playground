package repository

import (
	"fmt"
	"iwf-playground/model"
)

type OrderRepository interface {
	SetOrderStatus(orderId string, state string) error
	GetOrderById(orderId string) (model.Order, error)
}

type repository struct {
}

func NewOrderRepository() OrderRepository {
	return &repository{}
}

var mockRepo = []model.Order{
	{
		Id:          "order-id1",
		AssignedPPK: "ppk-1",
		State:       model.DRAFT,
	},
	{
		Id:          "order-id2",
		AssignedPPK: "ppk-2",
		State:       model.DRAFT,
	},
	{
		Id:          "order-id3",
		AssignedPPK: "ppk-3",
		State:       model.DRAFT,
	},
}

func (ths *repository) SetOrderStatus(orderId string, state string) error {
	idx := -1
	for i, order := range mockRepo {
		if order.Id == orderId {
			idx = i
			break
		}
	}

	if idx == -1 {
		return fmt.Errorf("no data")
	}

	currentOrder := mockRepo[idx]
	currentOrder.State = model.State(state)

	mockRepo[idx] = currentOrder

	return nil
}

func (ths *repository) GetOrderById(orderId string) (model.Order, error) {
	for _, order := range mockRepo {
		if order.Id == orderId {
			return order, nil
		}
	}

	return model.Order{}, fmt.Errorf("no data")
}
