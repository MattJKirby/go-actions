package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/app/config"
)

type baseExecutable struct {
	construct action.GoAction
	instance  *action.ActionInstance
}

func newBaseExecutable(config *config.GlobalConfig, typeDef *definition.ActionTypeDefinition) (*baseExecutable, error) {
	instance := action.NewActionInstance(typeDef.TypeName, config)
	defaultProps := typeDef.NewDefaultProps()
	ctor := typeDef.NewConstructor()
	action, err := ctor(instance, defaultProps)
	if err != nil {
		return nil, err
	}

	return &baseExecutable{
		construct: action,
		instance:  instance,
	}, nil
}
