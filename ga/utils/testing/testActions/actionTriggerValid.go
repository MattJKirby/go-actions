package testActions

import "go-actions/ga/action"

type ActionTriggerValid struct{
	*action.Internals
}

type ActionTriggerValidProps struct{}

func (atv ActionTriggerValid) Init(*action.ActionInstance) {}
func (atv ActionTriggerValid) Execute()                    {}
func (atv ActionTriggerValid) PublishTriggerConditions()   {}
