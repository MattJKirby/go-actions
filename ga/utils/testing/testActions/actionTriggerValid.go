package testActions

import "go-actions/ga/action"

type ActionTriggerValid struct{}

type ActionTriggerValidProps struct{}

func (atv ActionTriggerValid) Init(*action.ActionInstance) {}
func (atv ActionTriggerValid) Execute()                    {}
func (atv ActionTriggerValid) PublishTriggerConditions()   {}

func GenerateActionTriggerValidRegistration() action.GoActionRegistration[ActionTriggerValid] {
	return GenerateRegistration(ActionTriggerValid{})
}
