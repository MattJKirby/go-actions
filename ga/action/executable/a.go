package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
)

type ExecutableAction struct {
	instance *action.ActionInstance
}

func NewExecutableAction(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition) *ExecutableAction {
	return &ExecutableAction{
		instance: newExecutableInstance(modelConfig, typeDef, nil),
	}
}

func newExecutableInstance(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition, props action.GoActionProps) *action.ActionInstance {
	instance := action.NewActionInstance(typeDef.TypeName, modelConfig)
	ctor := typeDef.NewConstructor()

	if err := typeDef.ValidatePropsType(props); err != nil {
		ctor(instance, typeDef.NewDefaultProps())
		return instance
	}

	ctor(instance, props)
	return instance
}
