package execution

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

type InstanceResolver struct {
	app *app.App
	triggers map[string]*executable.Action[action.GoAction]
	actions map[string]*executable.Action[action.GoAction]
}

func NewInstanceResolver(app *app.App) *InstanceResolver {
	return &InstanceResolver{
		app: app,
		triggers: make(map[string]*executable.Action[action.GoAction]),
		actions: make(map[string]*executable.Action[action.GoAction]),
	}
}

func (ir *InstanceResolver) Resolve(flowDef flow.FlowDefinition) error{
	for _, instance := range flowDef.Actions.Store.GetEntries() {
		_, err := flowDef.NewAction(instance.Name)
		if err != nil {
			return err
		}


	}

	return nil
}