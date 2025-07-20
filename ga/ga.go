package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction](act T, reg *action.ActionRegistration[T]) {
	app.RegisterAction(act, reg)(ga)
}

func GetRegisteredTypeDefinition[T action.GoAction](act T) (*action.TypeDefinition, error) {
	return app.GetDefinitionByType(act)(ga)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

func AddAction[T action.GoAction](f *flow.Flow, a T) *action.Action[T] {
	act, _ := flow.AddAction(f, func(a *action.Action[T]) {})
	return act
}

func AddActionConfigurable[T action.GoAction](f *flow.Flow, fn func(*action.Action[T])) *action.Action[T] {
	act, _ := flow.AddAction(f, fn)
	return act
}
