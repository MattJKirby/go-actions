package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
	"go-actions/ga/app/config"
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
		ActionInput:  model.Input(inst.Model, "ActionInput", false, nil),
		ActionOutput: model.Output(inst.Model, "ActionOutput", nil),
	}
}

func NewAction[T action.GoAction](config *config.GlobalConfig, typeDef *definition.ActionTypeDefinition) (*Action[T], error) {
	instance := action.NewActionInstance(typeDef.TypeName, config)

	action, err := definition.NewAction[T](typeDef, instance, nil)
	if err != nil {
		return nil, err
	}

	return &Action[T]{
		BaseActionFields: NewBaseActionFields(instance),
		Definition:       action,
		Instance:         instance,
	}, nil
}
