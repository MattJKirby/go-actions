package examples

import (
	"go-actions/ga"
	"go-actions/ga/action"
)

func init() {
	ga.RegisterAction[*BasicAction](nil)
}

type BasicAction struct {}

func (ba *BasicAction) Init(inst *action.ActionInstance){}

func (ba *BasicAction) Execute(){}