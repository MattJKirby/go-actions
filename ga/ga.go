package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
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

func GetAction[T action.GoAction, P action.GoActionProps]() (*app.InstantiatedTypedAction[T], error) {
	instantiatedAction, err := app.InstantiateActionFromType[T, P](nil)(ga)
	if err != nil {
		return nil, err
	}
	return	instantiatedAction, nil
}

func NewFlowAction[T action.GoAction, P action.GoActionProps](f *flow.Flow, props *P) *T {
	action, err := flow.AddAction[T](f, props)
	if err != nil {
		panic("could not retrieve action from app")
	}
	return action.Action
}

func NewFlow() *flow.Flow {
	return flow.NewFlow(ga)
}

// func GetActionByName(name string)
