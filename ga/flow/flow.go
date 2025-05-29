package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
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

func AddAction[T action.GoAction](f *Flow) (*executable.Action[T], error) {
	typeDef, err :=  app.GetDefinitionByType[T]()(f.flowApp)
	if err != nil {
		return nil, err
	}
	return addAction[T](f.Definition, typeDef)

}