package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/action/model/io"
	"go-actions/ga/app"
	"go-actions/ga/libs/store"
)

type flowDefinition struct {
	app *app.App
	Actions *store.BaseStore[action.ActionInstance] `json:"Actions"`
}

func NewFlowDefinition(app *app.App) *flowDefinition {
	return &flowDefinition{
		app: app,
		Actions: store.NewBaseStore[action.ActionInstance](),
	}
}


func (fd *flowDefinition) NewAction(actionName string) (*executable.Action[action.GoAction], error) {
	action, err := app.GetActionByName(actionName)(fd.app)
	if err != nil {
		return nil, err
	}

	fd.Actions.Insert(action.Instance.Model.ActionUid, action.Instance)
	return action, nil
}

func (fd *flowDefinition) NewReference(sourceActionUid string, sourceId string, targetActionUid string, targetId string) error {
	sourceAction, err := fd.Actions.Get(sourceActionUid)
	if err != nil {
		return err
	}

	source, err := sourceAction.Model.Outputs.Get(sourceId)
	if err != nil {
		return err
	}

	targetAction, err := fd.Actions.Get(targetActionUid)
	if err != nil {
		return err
	}

	target, err := targetAction.Model.Inputs.Get(targetId)
	if err != nil {
		return err
	}

	if err := io.NewActionReference(fd.app.Config.Global, source, target).AssignReferences(); err != nil {
		return err
	}
	return nil
}
