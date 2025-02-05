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