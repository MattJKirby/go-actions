package flow

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp        *app.App
	flowDefinition *flowDefinition
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:        app,
		flowDefinition: NewFlowDefinition(),
	}
}

func NewFlowAction[T action.GoAction, P action.GoActionProps](f *Flow, props *P) (*executable.BaseExecutable[T], error) {
	if instantiated, err := app.GetAction[T](props)(f.flowApp); err == nil {
		f.flowDefinition.AddInstance(instantiated.Instance)
		return instantiated, nil
	}

	return nil, fmt.Errorf("could not retrieve action from app")
}
