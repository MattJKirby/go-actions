package executable

import (
	"fmt"
	"go-actions/ga/action"

	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
)

type Action[T action.GoAction] struct {
	*BaseActionFields
	Definition T
	Instance   *action.ActionInstance
	typeDef    *action.TypeDefinition
}

type BaseActionFields struct {
	ActionInput  *input.ActionInput
	ActionOutput *output.ActionOutput
}

func NewBaseActionFields(inst *action.ActionInstance) *BaseActionFields {
	return &BaseActionFields{
		ActionInput:  model.Input(inst.Model, "Action", false, nil),
		ActionOutput: model.Output(inst.Model, "Action", nil),
	}
}

func NewAction[T action.GoAction](typeDef *action.TypeDefinition, inst *action.ActionInstance) (*Action[T], error) {
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
