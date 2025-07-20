package action

import (
	"fmt"

	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
)

type Action[T GoAction] struct {
	*BaseActionFields
	Definition T
	Instance   *ActionInstance
	typeDef    *TypeDefinition
}

type BaseActionFields struct {
	ActionInput  *io.ActionInput
	ActionOutput *io.ActionOutput
}

func NewBaseActionFields(inst *ActionInstance) *BaseActionFields {
	return &BaseActionFields{
		ActionInput:  model.Input(inst.Model, "Action", false, nil),
		ActionOutput: model.Output(inst.Model, "Action", nil),
	}
}

func NewAction[T GoAction](typeDef *TypeDefinition, inst *ActionInstance) (*Action[T], error) {
	act, ok := typeDef.ActionValue.Interface().(T)
	if !ok {
		return nil, fmt.Errorf("new action does not match expected type")
	}

	act.Init(inst)

	return &Action[T]{
		BaseActionFields: NewBaseActionFields(inst),
		Instance:         inst,
		Definition:       act,
		typeDef:          typeDef,
	}, nil
}
