package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp    *app.App
	Definition *FlowDefinition
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:    app,
		Definition: NewFlowDefinition(app),
	}
}

func AddAction[T action.GoAction](f *Flow, configurationFn func(*action.Action[T])) (*action.Action[T], error) {
	actionType := new(T)
	typeDef, err := app.GetDefinitionByType(*actionType)(f.flowApp)
	if err != nil {
		return nil, err
	}

	action, err := addAction[T](f.Definition, typeDef)
	if err != nil {
		return nil, err
	}

	configurationFn(action)
	return action, nil
}
