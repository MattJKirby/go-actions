package action

import (
	"go-actions/ga/utils"
	"reflect"
)

type ActionDefinition struct {
	Name     string
	TypePath string
	ActionTypeDefinition
}

func NewActionDefinition[T Action](def Constructor[T]) (*ActionDefinition, error) {
	typeDef := TypeDefinitionFromConstructor(def)

	return &ActionDefinition{
		Name:     utils.TypeName(typeDef.actionType),
		TypePath: utils.TypePath(typeDef.actionType),
		ActionTypeDefinition: *typeDef,
	}, nil
}

func (ad *ActionDefinition) ActionType() reflect.Type {
	return ad.actionType
}

func (ad *ActionDefinition) Constructor() reflect.Value {
	return ad.ctorValue
}