package flow

import "go-actions/ga/action"

type FlowDefinition struct {
	actions map[string]*action.ActionInstance
}

func NewFlowDefinition() *FlowDefinition {
	return &FlowDefinition{
		actions: make(map[string]*action.ActionInstance),
	}
}

func (fd *FlowDefinition) AddInstance(instance *action.ActionInstance){
	fd.actions[instance.Model.ActionUid] = instance
}