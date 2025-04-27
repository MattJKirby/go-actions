package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction](registration *action.GoActionRegistration[T]) {
	app.RegisterAction(registration)(ga)
}

func NewAction[T action.GoAction](f *flow.Flow) (*executable.Action[T], error) {
	return flow.NewFlowAction[T](f)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
