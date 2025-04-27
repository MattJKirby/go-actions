package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction](reg *action.ActionRegistration[T]) {
	app.RegisterAction(reg)(ga)
}

func GetRegisteredTypeDefinition[T action.GoAction]() (*definition.ActionTypeDefinition, error) {
	return app.GetDefinitionByType[T]()(ga)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
