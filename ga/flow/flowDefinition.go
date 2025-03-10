package flow

import "go-actions/ga/action"

type flowDefinition struct {
	Actions map[string]*action.ActionInstance `json:"Actions"`
}

func NewFlowDefinition() *flowDefinition {
	return &flowDefinition{
		Actions: make(map[string]*action.ActionInstance),
	}
}

func (fd *flowDefinition) AddInstance(instance *action.ActionInstance) {
	fd.Actions[instance.Model.ActionUid] = instance
}
