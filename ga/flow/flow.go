package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp         *app.App
	FlowDefinition []*action.ActionInstance `json:"actions"`
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:         app,
		FlowDefinition: make([]*action.ActionInstance, 0),
	}
}

func AddAction[T action.GoAction, P action.GoActionProps](f *Flow, props *P) (*app.InstantiatedTypedAction[T], error) {
	instantiated, err := app.InstantiateActionFromType[T](props)(f.flowApp)
	if err != nil {
		return nil, err
	}

	f.FlowDefinition = append(f.FlowDefinition, instantiated.Instance)
	return instantiated, err
}
