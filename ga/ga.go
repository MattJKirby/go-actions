package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction, Props action.GoActionProps](registration *action.GoActionRegistration[T, Props]) {
	app.RegisterAction(registration)(ga)
}

func GetAction[T action.GoAction, P action.GoActionProps]() (*executable.BaseExecutable[T], error) {
	return app.GetAction[T, P](nil)(ga)
}

func NewFlowAction[T action.GoAction, P action.GoActionProps](f *flow.Flow, props *P) (T, error) {
	act, err := flow.NewFlowAction[T](f, props)
	return act.Action, err
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga, flow.NewFlowDefinition())
}

// func GetActionByName(name string)
