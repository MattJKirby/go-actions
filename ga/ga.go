package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
	"go-actions/ga/flow"
)

var ga = app.NewApp("GoActionsDefaultApp")

func RegisterAction[T action.GoAction, Props action.GoActionProps](registration *action.GoActionRegistration[T, Props]) {
	app.RegisterAction(registration)(ga)
}

func GetActionRegistration[T action.GoAction, P action.GoActionProps]() (*definition.ActionDefinition[T, P], error) {
	return app.GetActionRegistration[T, P]()(ga)
}

func GetAction[T action.GoAction, P action.GoActionProps]() (*executable.TypedExecutable[T, P], error) {
	return app.InstantiateTypedAction[T, P](nil)(ga)
}

func NewFlowAction[T action.GoAction, P action.GoActionProps](f *flow.Flow, props *P) *executable.TypedExecutable[T, P] {
	action, err := flow.AddAction[T](f, props)
	if err != nil {
		panic("could not retrieve action from app")
	}
	return action
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
