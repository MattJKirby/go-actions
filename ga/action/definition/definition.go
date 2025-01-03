package definition

import (
	"go-actions/ga/action"
	"go-actions/ga/utils"
)

type ActionDefinition[T action.GoAction, P action.GoActionProps] struct {
	Registration   *action.GoActionRegistration[T, P]
	TypeDefinition *ActionTypeDefinition
	Name           string
	TypePath       string
}

func NewActionDefinition[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) (*ActionDefinition[T, P], error) {
	typeDef := TypeDefinitionFromConstructor(reg.Constructor)

	return &ActionDefinition[T, P]{
		Registration:   reg,
		TypeDefinition: typeDef,
		Name:           utils.TypeName(typeDef.ActionType),
		TypePath:       utils.TypePath(typeDef.ActionType),
	}, nil
}
