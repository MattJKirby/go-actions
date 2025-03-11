package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
)

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

func (fd *flowDefinition) NewAction(flowApp *app.App, actionName string) (*app.InitialisedAction, error) {
	action, err := app.GetActionByName(actionName)(flowApp)
	if err != nil {
		return nil, err
	}
	
	fd.AddInstance(action.InitialisedInstance)
	return action, nil
}
