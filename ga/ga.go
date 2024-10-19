package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
)

var ga = app.NewApp()

func DefineAction[T action.Action](actionConstructor action.Constructor[T]) *action.ActionDefinition {
	return app.RegisterAction(actionConstructor)(ga)
}

func GetActionDefinition[T action.Action](action T) (*action.ActionDefinition, error) {
	return ga.GetActionDef(action)
}

func GetAction[T action.Action](a T) (*action.GoAction[T], error) {
	return app.NewAction[T](a)(ga)
}

// func GetActionByName(name string)
