package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
	"go-actions/ga/app/config"
)

type Action[T action.GoAction] struct {
	Definition   T
	Instance *action.ActionInstance
}

type baseFields struct {
	ActionInput *input.ActionInput
  ActionOutput *output.ActionOutput
}

func NewAction[T action.GoAction](config *config.GlobalConfig, typeDef *definition.ActionTypeDefinition) (*Action[T], error) {
	instance := action.NewActionInstance(typeDef.TypeName, config)

	action, err := definition.NewAction[T](typeDef, instance, nil)
	if err != nil {
		return nil, err
	}

	return &Action[T]{
		Definition:   action,
		Instance: instance,
	}, nil
}
