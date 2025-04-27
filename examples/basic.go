package examples

import (
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/flow"
)

func init() {
	ga.RegisterAction(&action.GoActionRegistration[*BasicAction]{
		Action:       &BasicAction{},
	})
}

func NewBasicAction(flow *flow.Flow) (*executable.Action[*BasicAction], error) {
	return ga.NewAction[*BasicAction](flow)
}

type BasicAction struct {}

func (ba *BasicAction) Init(inst *action.ActionInstance){}

func (ba *BasicAction) Execute(){}