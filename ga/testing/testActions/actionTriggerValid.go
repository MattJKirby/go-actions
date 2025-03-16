package testActions

import "go-actions/ga/action"

type ActionTriggerValid struct {}

type ActionTriggerValidProps struct {}

func (atv ActionTriggerValid) Execute(){

}

func newActionTriggerValid(instance *action.ActionInstance, props ActionTriggerValidProps) *ActionTriggerValid {
	return &ActionTriggerValid{}
}

func GenerateActionTriggerValidRegistration() action.GoActionRegistration[ActionTriggerValid, ActionTriggerValidProps]{
	return GenerateRegistration(newActionTriggerValid, &ActionTriggerValidProps{})
}