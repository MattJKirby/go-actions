package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
)

type ExecutableAction struct {
	instance *action.ActionInstance
}

func NewExecutableAction(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition, props action.GoActionProps) *ExecutableAction {
	return &ExecutableAction{
		instance: newExecutableInstance(modelConfig, typeDef),
	}
}

func newExecutableInstance(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition) *action.ActionInstance {
	instance := action.NewActionInstance(typeDef.TypeName, modelConfig)
	defaultProps := typeDef.NewDefaultProps()
	ctor := typeDef.NewConstructor()

	ctor(instance, defaultProps)
	return instance
}
