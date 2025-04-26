package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/app/config"
)

type Action[T action.GoAction] struct {
	Action   T
	Instance *action.ActionInstance
}

func NewAction[T action.GoAction](config *config.GlobalConfig, typeDef *definition.ActionTypeDefinition) (*Action[T], error) {
	instance := action.NewActionInstance(typeDef.TypeName, config)

	action, err := definition.NewAction[T](typeDef, instance, nil)
	if err != nil {
		return nil, err
	}

	return &Action[T]{
		Action:   action,
		Instance: instance,
	}, nil
}
