package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
)

var ga = app.NewApp()

func DefineAction[T action.GoAction](actionConstructor action.GoActionConstructor[T]) *definition.ActionDefinition {
	return app.RegisterAction(actionConstructor)(ga)
}

func GetActionDefinition[T action.GoAction](action T) (*definition.ActionDefinition, error) {
	return ga.GetActionDef(action)
}

func GetAction[T action.GoAction](a T) (*executable.Action[T], error) {
	return app.NewAction[T](a)(ga)
}

// func GetActionByName(name string)
