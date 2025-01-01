package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/app/registration"
	"go-actions/ga/flow"
)

var ga = app.NewApp()

func DefineAction[T action.GoAction](registration *action.GoActionRegistration[T]) *registration.RegisteredAction[T] {
	return app.RegisterAction(registration)(ga)
}

func GetActionRegistration[T action.GoAction](action T) (*registration.RegisteredAction[T], error) {
	return app.GetActionRegistration[T](action)(ga)
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
