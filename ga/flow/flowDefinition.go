package flow

import "go-actions/ga/action"

type FlowDefinition struct {
	Actions map[string]*action.ActionInstance `json:"Actions"`
}

func NewFlowDefinition() *FlowDefinition {
	return &FlowDefinition{
		Actions: make(map[string]*action.ActionInstance),
	}
}

func (fd *FlowDefinition) AddInstance(instance *action.ActionInstance) {
	fd.Actions[instance.Model.ActionUid] = instance
}


