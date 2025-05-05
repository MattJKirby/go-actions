package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/action/model"
	"go-actions/ga/app"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type flowDefinition struct {
	app     *app.App
	Actions *store.ResourceStore[action.ActionInstance] `json:"actions"`
}

func NewFlowDefinition(app *app.App) *flowDefinition {
	return &flowDefinition{
		app:     app,
		Actions: store.NewResourceStore[action.ActionInstance](false),
	}
}

func (fd *flowDefinition) NewAction(actionName string) (*executable.Action[action.GoAction], error) {
	action, err := app.GetActionByName(actionName)(fd.app)
	if err != nil {
		return nil, err
	}

	fd.Actions.NewResource(*action.Instance)
	return action, nil
}

func (fd *flowDefinition) NewReference(sourceUid *uid.ResourceUid, targetUid *uid.ResourceUid) error {
	sourceAction, err := fd.Actions.GetResource(sourceUid.GetBaseUid())
	if err != nil {
		return err
	}
	source, err := sourceAction.Model.Outputs.GetResource(sourceUid.GetUid())
	if err != nil {
		return err
	}
	targetAction, err := fd.Actions.GetResource(targetUid.GetBaseUid())
	if err != nil {
		return err
	}
	target, err := targetAction.Model.Inputs.GetResource(targetUid.GetUid())
	if err != nil {
		return err
	}
	model.Reference(fd.app.Config.Global, source, target)
	return nil
}
