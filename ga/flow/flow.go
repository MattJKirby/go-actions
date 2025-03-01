package flow

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp        *app.App
	FlowDefinition []*action.ActionInstance `json:"actions"`
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:        app,
		FlowDefinition: make([]*action.ActionInstance, 0),
	}
}

func NewFlowAction[T action.GoAction, P action.GoActionProps](f *Flow, props *P) (*app.InitialisedTypedAction[T], error) {
	instantiated, err := app.GetAction[T](props)(f.flowApp)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve action from app")
	}

	f.FlowDefinition = append(f.FlowDefinition, instantiated.InitialisedInstance)
	return instantiated, nil
}
