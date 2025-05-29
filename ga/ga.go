package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction](reg *action.ActionRegistration[T]) {
	app.RegisterAction(reg)(ga)
}

func GetRegisteredTypeDefinition[T action.GoAction]() (*action.TypeDefinition, error) {
	return app.GetDefinitionByType[T]()(ga)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

func AddAction[T action.GoAction](f *flow.Flow) (*executable.Action[T]) {
	act, _ := flow.AddAction[T](f)
	return act
}

