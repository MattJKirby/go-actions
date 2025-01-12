package testActions

import (
	"go-actions/ga/action"
	"go-actions/ga/action/model/parameter"
)

type ActionValid struct{
	Param1 *parameter.ActionParameter[string]
}

type ActionValidProps struct{
	Param1 string
}

func (tav ActionValid) Execute() {}

func newActionValid(instance *action.ActionInstance, props ActionValidProps) *ActionValid {
	return &ActionValid{
		Param1: action.Parameter(instance, "param1", props.Param1),
	}
}

func GenerateActionValidRegistration() action.GoActionRegistration[ActionValid, ActionValidProps] {
	return GenerateRegistration(newActionValid, &ActionValidProps{})
}
