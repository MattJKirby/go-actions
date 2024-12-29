package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp()

func RegisterAction[T action.GoAction, Props any](registration *action.GoActionRegistration[T, Props]) *definition.ActionDefinition {
	return app.RegisterAction(registration)(ga)
}

func GetActionDefinition[T action.GoAction](action T) (*definition.ActionDefinition, error) {
	return app.GetActionDefinition(action)(ga)
}

func GetAction[T action.GoAction, Props any](a T, props *Props) (*executable.Action[T, Props], error) {
	return app.GetAction[T, Props](a)(ga)
}

func NewAction[T action.GoAction, Props any](f *flow.Flow, a T, props *Props) *executable.Action[T, Props] {
	return flow.AddAction(a, props)(f)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
