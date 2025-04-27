package examples

import (
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/flow"
)

func init() {
	ga.RegisterAction(&action.GoActionRegistration[*BasicAction, BasicActionProps]{
		Action:       &BasicAction{},
		DefaultProps: BasicActionProps{},
	})
}

func NewBasicAction(flow *flow.Flow, props *BasicActionProps) (*executable.Action[*BasicAction], error) {
	return ga.NewAction[*BasicAction](flow, props)
}

type BasicAction struct {}
type BasicActionProps struct {}

func (ba *BasicAction) Init(inst *action.ActionInstance){}

func (ba *BasicAction) Execute(){}