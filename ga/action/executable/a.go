package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
)

type ExecutableAction struct {
	instance *action.ActionInstance
	Action *action.GoAction
}

func NewExecutableAction(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition) *ExecutableAction {
	instance, action, _ := newExecutableInstance(modelConfig, typeDef, nil)
	
	return &ExecutableAction{
		instance,
		action,
	}
}

func newExecutableInstance(modelConfig model.ActionModelConfig, typeDef *definition.ActionTypeDefinition, props action.GoActionProps) (*action.ActionInstance, *action.GoAction, error) {
	instance := action.NewActionInstance(typeDef.TypeName, modelConfig)
	ctor := typeDef.NewConstructor()

	if props == nil {
		props = typeDef.NewDefaultProps()
	}

	action , err := ctor(instance, props)
	if err != nil {
		return nil, nil, err
	}
	
	return instance, &action, nil
}
