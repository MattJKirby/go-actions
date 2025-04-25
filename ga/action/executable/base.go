package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/app/config"
)

type BaseExecutable[T action.GoAction] struct {
	Action   T
	Instance *action.ActionInstance
}

func NewBaseExecutable[T action.GoAction](config *config.GlobalConfig, typeDef *definition.ActionTypeDefinition) (*BaseExecutable[T], error) {
	instance := action.NewActionInstance(typeDef.TypeName, config)

	action, err := definition.NewAction[T](typeDef, instance, nil)
	if err != nil {
		return nil, err
	}

	return &BaseExecutable[T]{
		Action:   action,
		Instance: instance,
	}, nil
}
