package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
)

type ExecutableAction struct {

}

func NewExecutableAction(typeDef *definition.ActionTypeDefinition, props *action.GoAction) *ExecutableAction {
	return &ExecutableAction{}
}

func newExecutableInstance(typeDef *definition.ActionTypeDefinition) *action.ActionInstance{
	instance := action.NewActionInstance(typeDef.TypeName)
	props := typeDef.NewDefaultProps()
	ctor := typeDef.NewConstructor()

	ctor(instance, props)
	return instance
}