package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/action/model"
	"go-actions/ga/app"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type FlowDefinition struct {
	app     *app.App
	Actions *store.ResourceStore[action.ActionInstance] `json:"actions"`
}

func NewFlowDefinition(app *app.App) *FlowDefinition {
	return &FlowDefinition{
		app:     app,
		Actions: store.NewResourceStore[action.ActionInstance](false),
	}
}

func (fd *FlowDefinition) NewAction(actionName string) (*executable.Action[action.GoAction], error) {
	typeDef, err := app.GetDefinitionByName(actionName)(fd.app)
	if err != nil {
		return nil, err
	}
	
	action, err := app.GetAction(typeDef, nil)(fd.app)
	if err != nil {
		return nil, err
	}

	fd.Actions.NewResource(*action.Instance)
	return action, nil
}

func (fd *FlowDefinition) NewReference(sourceUid uid.ResourceUid, targetUid uid.ResourceUid) error {
	sourceAction, err := fd.Actions.GetResource(sourceUid.BaseUid())
	if err != nil {
		return err
	}
	source, err := sourceAction.Model.Outputs.GetResource(sourceUid.FullUid())
	if err != nil {
		return err
	}
	targetAction, err := fd.Actions.GetResource(targetUid.BaseUid())
	if err != nil {
		return err
	}
	target, err := targetAction.Model.Inputs.GetResource(targetUid.FullUid())
	if err != nil {
		return err
	}
	model.Reference(fd.app.Config.Global, source, target)
	return nil
}
