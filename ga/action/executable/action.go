package executable

import (
	"go-actions/ga/action"

	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
)

type Action[T action.GoAction] struct {
	*BaseActionFields
	Definition T
	Instance   *action.ActionInstance
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
	action, err := action.InitialiseInstance[T](typeDef, inst)
	if err != nil {
		return nil, err
	}

	return &Action[T]{
		BaseActionFields: NewBaseActionFields(inst),
		Definition:       action,
		Instance:         inst,
	}, nil
}