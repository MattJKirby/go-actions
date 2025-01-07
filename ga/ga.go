package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp()

func RegisterAction[T action.GoAction, Props action.GoActionProps](registration *action.GoActionRegistration[T, Props]) {
	app.RegisterAction(registration)(ga)
}

func GetActionRegistration[T action.GoAction](a T) (*definition.ActionDefinition[T, action.GoActionProps], error) {
	return app.GetActionRegistration[T, action.GoActionProps](a)(ga)
}

func GetAction[T action.GoAction](a T) (*executable.Action[T, action.GoActionProps], error) {
	return app.GetAction[T, action.GoActionProps](a, nil)(ga)
}

func ActionFunction[T action.GoAction, P action.GoActionProps](f *flow.Flow, a T, props *P) *executable.Action[T, P] {
	return flow.AddAction(a, props)(f)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
