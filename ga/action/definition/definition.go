package definition

import (
	"go-actions/ga/action"
	"go-actions/ga/utils"
)

type ActionDefinition struct {
	Name     string
	TypePath string
	ActionTypeDefinition
}

func NewActionDefinition[T action.GoAction](def action.GoActionConstructor[T]) (*ActionDefinition, error) {
	typeDef := TypeDefinitionFromConstructor(def)

	return &ActionDefinition{
		Name:                 utils.TypeName(typeDef.ActionType),
		TypePath:             utils.TypePath(typeDef.ActionType),
		ActionTypeDefinition: *typeDef,
	}, nil
}
