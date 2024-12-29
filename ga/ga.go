package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp()

func DefineAction[T action.GoAction](actionConstructor action.GoActionConstructor[T]) *definition.ActionDefinition {
	return app.DefineAction(actionConstructor)(ga)
}

func GetActionDefinition[T action.GoAction](action T) (*definition.ActionDefinition, error) {
	return app.GetActionDefinition(action)(ga)
}

func GetAction[T action.GoAction](a T) (*executable.Action[T], error) {
	return app.GetAction[T](a)(ga)
}

func NewAction[T action.GoAction](f *flow.Flow, a T) *executable.Action[T] {
	return flow.AddAction(a)(f)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
