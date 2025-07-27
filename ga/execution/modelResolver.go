package execution

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

type InstanceResolver struct {
	app      *app.App
	triggers map[string]*action.Action[action.GoAction]
	actions  map[string]*action.Action[action.GoAction]
	sources  map[string]struct{}
}

func NewInstanceResolver(app *app.App) *InstanceResolver {
	return &InstanceResolver{
		app:      app,
		triggers: make(map[string]*action.Action[action.GoAction]),
		actions:  make(map[string]*action.Action[action.GoAction]),
		sources:  make(map[string]struct{}),
	}
}

func (ir *InstanceResolver) Resolve(flowDef flow.FlowDefinition) error {
	for id, instance := range flowDef.Actions.Store.GetEntries() {
		typeDef, err := app.GetDefinitionByName(instance.Name)(ir.app)
		if err != nil {
			return err
		}

		act, err := app.GetAction[action.GoAction](typeDef, &instance)(ir.app)
		if err != nil {
			return err
		}

		if ir.isSource(act) {
			ir.sources[id] = struct{}{}
		}

		ir.actions[id] = act
	}
	return nil
}

func (ir *InstanceResolver) isSource(act *action.Action[action.GoAction]) bool {
	if len(act.Instance.Model.Inputs.Store.GetEntries()) != 1 {
		return false
	}

	if len(act.Definition.GetInput().SourceReferences.Store.GetEntries()) != 0 {
		return false
	}
	return true
}
