package action

import (
	"go-actions/ga/utils"
)

type ActionDefinition struct {
	Name     string
	TypePath string
	ActionTypeDefinition
}

func NewActionDefinition[T GoAction](def GoActionConstructor[T]) (*ActionDefinition, error) {
	typeDef := TypeDefinitionFromConstructor(def)

	return &ActionDefinition{
		Name:     utils.TypeName(typeDef.ActionType),
		TypePath: utils.TypePath(typeDef.ActionType),
		ActionTypeDefinition: *typeDef,
	}, nil
}