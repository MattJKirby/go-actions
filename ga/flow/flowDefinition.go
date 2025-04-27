package flow

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
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

func (fd *flowDefinition) addInstance(instance *action.ActionInstance) error {
	if _,exists := fd.Actions[instance.Model.ActionUid]; !exists {
		fd.Actions[instance.Model.ActionUid] = instance
		return nil
	}
	return fmt.Errorf("error adding instance: instance '%s' already exists", instance.Model.ActionUid)
}

func (fd *flowDefinition) NewAction(flowApp *app.App, actionName string) (*executable.Action[action.GoAction], error) {
	action, err := app.GetActionByName(actionName)(flowApp)
	if err != nil {
		return nil, err
	}

	fd.addInstance(action.Instance)
	return action, nil
}
