package definition

import (
	"go-actions/ga/action"
)

type ActionDefinition[T action.GoAction, P action.GoActionProps] struct {
	*action.GoActionRegistration[T, P]
	*ActionTypeDefinition
}

func NewActionDefinition[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) *ActionDefinition[T, P] {
	return &ActionDefinition[T, P]{
		GoActionRegistration: reg,
		ActionTypeDefinition: TypeDefinitionFromRegistration(reg),
	}
}

func (ad *ActionDefinition[T, P]) GetTypeDefinition() *ActionTypeDefinition {
	return ad.ActionTypeDefinition
}
