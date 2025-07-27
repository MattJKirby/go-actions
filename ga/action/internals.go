package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
)

type ActionInternals struct {
	instance *ActionInstance
	actionInput  *io.ActionInput
	actionOutput *io.ActionOutput
}

func NewActionInternals(instance *ActionInstance) *ActionInternals {
	return &ActionInternals{
		instance: instance,
		actionInput:  model.Input(instance.Model, "Action", false, nil),
		actionOutput: model.Output(instance.Model, "Action", nil),
	}
}

func (ai *ActionInternals) GetInstance() *ActionInstance {
	return ai.instance
}