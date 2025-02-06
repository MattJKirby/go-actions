package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
)

type ExecutableAction struct {

}

func NewExecutableAction(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition, props *action.GoAction) *ExecutableAction {
	return &ExecutableAction{}
}

func newExecutableInstance(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition) *action.ActionInstance{
	instance := action.NewActionInstance(typeDef.TypeName, modelConfig)
	props := typeDef.NewDefaultProps()
	ctor := typeDef.NewConstructor()

	ctor(instance, props)
	return instance
}