package create_order

import (
	"iwf-playground/repository"
	"iwf-playground/service/workflow/create_order/state"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type createOrderWorkflow struct {
	iwf.DefaultWorkflowType

	repo repository.OrderRepository
}

func NewCreateOrderWorkflow(repo repository.OrderRepository) *createOrderWorkflow {
	return &createOrderWorkflow{
		repo: repo,
	}
}

func (ths *createOrderWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(state.NewInitState()),
		iwf.NonStartingStateDef(state.NewSetOrderToReview(ths.repo)),
		iwf.NonStartingStateDef(state.NewSendNotificationState(ths.repo)),
		iwf.NonStartingStateDef(state.NewExpiredTimerState()),
		iwf.NonStartingStateDef(state.NewSignalOrderPPKState()),
		iwf.NonStartingStateDef(state.NewSetOrderToCancelled(ths.repo)),
		iwf.NonStartingStateDef(state.NewSetOrderToRejected(ths.repo)),
		iwf.NonStartingStateDef(state.NewSetOrderToCompleted(ths.repo)),
	}
}

func (ths *createOrderWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return []iwf.PersistenceFieldDef{
		iwf.DataObjectDef(state.KeyOrderId),
	}
}

func (ths *createOrderWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.SignalChannelDef(state.PPK_CHANNEL),
	}
}
