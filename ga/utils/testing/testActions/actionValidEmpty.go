package testActions

import (
	"go-actions/ga/action"
)

type ActionValidEmpty struct{
	*action.Internals
}

type ActionValidEmptyProps struct{}

func (tav ActionValidEmpty) Init(*action.ActionInstance) {}
func (tav ActionValidEmpty) Execute()                    {}
