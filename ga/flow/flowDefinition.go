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

func (fd *flowDefinition) GetModels() []*model.ActionModel {
	models := make([]*model.ActionModel, 0, len(fd.Actions.Store.GetEntries()))
	for _, action := range fd.Actions.Store.GetEntries() {
		models = append(models, action.Model)
	}
	return models
}

func (fd *flowDefinition) NewAction(actionName string) (*executable.Action[action.GoAction], error) {
	action, err := app.GetActionByName(actionName)(fd.app)
	if err != nil {
		return nil, err
	}

	fd.Actions.NewResource(*action.Instance)
	return action, nil
}

func (fd *flowDefinition) NewReference(sourceUid uid.ResourceUid, targetUid uid.ResourceUid) error {
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
