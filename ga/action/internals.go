package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
)

type ActionInternals interface {
	InitInternals(*Internals)
	GetInstance() *ActionInstance
	GetOutput() *io.ActionOutput
	GetInput() *io.ActionInput
}

type Internals struct {
	instance *ActionInstance
	actionInput  *io.ActionInput
	actionOutput *io.ActionOutput
}

func NewInternals(instance *ActionInstance) *Internals {
	return &Internals{
		instance: instance,
		actionInput:  model.Input(instance.Model, "Action", false, nil),
		actionOutput: model.Output(instance.Model, "Action", nil),
	}
}

func (ai *Internals) GetInstance() *ActionInstance {
	return ai.instance
}

func (ai *Internals) SetInstance(i *ActionInstance) {
	ai.instance = i
}

func (ai *Internals) GetOutput() *io.ActionOutput {
	return ai.actionOutput
}

func (ai *Internals) GetInput() *io.ActionInput {
	return ai.actionInput
}

func (ai *Internals) InitInternals(i *Internals) {
	*ai = *i 
}	