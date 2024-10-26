package action

import (
	"go-actions/ga/utils"
	"reflect"
)

type ActionDefinition struct {
	name     string
	typePath string
	ActionTypeDefinition
}

func NewActionDefinition[T Action](def Constructor[T]) (*ActionDefinition, error) {
	typeDef := TypeDefinitionFromConstructor(def)

	return &ActionDefinition{
		name:     utils.TypeName(typeDef.actionType),
		typePath: utils.TypePath(typeDef.actionType),
		ActionTypeDefinition: *typeDef,
	}, nil
}

func (ad *ActionDefinition) ActionType() reflect.Type {
	return ad.actionType
}

func (ad *ActionDefinition) Constructor() reflect.Value {
	return ad.ctorValue
}

func (ad *ActionDefinition) Name() string {
	return ad.name
}

func (ad *ActionDefinition) TypeName() string {
	return ad.typePath
}
