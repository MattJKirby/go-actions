package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction](registration *action.GoActionRegistration[T]) {
	app.RegisterAction(registration)(ga)
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
