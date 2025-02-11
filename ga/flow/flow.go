package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp         *app.App
	ActionInstances []*action.ActionInstance `json:"actions"`
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:         app,
		ActionInstances: make([]*action.ActionInstance, 0),
	}
}

func AddAction[T action.GoAction, P action.GoActionProps](f *Flow, props *P) (*executable.TypedExecutable[T, P], error) {
	act, err := app.InstantiateTypedAction[T](props)(f.flowApp)
	if err != nil {
		return nil, err
	}

	f.ActionInstances = append(f.ActionInstances, act.Instance)
	return act, err
}
