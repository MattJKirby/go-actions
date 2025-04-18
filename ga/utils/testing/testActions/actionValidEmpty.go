package testActions

import (
	"go-actions/ga/action"
)

type ActionValidEmpty struct{}

type ActionValidEmptyProps struct{}

func (tav ActionValidEmpty) Init(*action.ActionInstance) {}
func (tav ActionValidEmpty) Execute()                    {}

func GenerateActionValidEmptyRegistration() action.GoActionRegistration[ActionValidEmpty, ActionValidEmptyProps] {
	return GenerateRegistration(ActionValidEmpty{}, ActionValidEmptyProps{})
}
