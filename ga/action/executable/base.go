package executable

import (
	"fmt"
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

	rawAction, err := typeDef.NewAction()
	if err != nil {
		return nil, err
	}

	act, ok := any(rawAction).(T)
	if !ok {
		return nil, fmt.Errorf("type assertion to generic T failed")
	}

	act.Init(instance)

	return &BaseExecutable[T]{
		Action:   act,
		Instance: instance,
	}, nil
}