package action

import (
	"fmt"
	"go-actions/ga/utils"
	"reflect"
)

type ActionTypeDefinition struct {
	ctorValue    reflect.Value
	ctorType    reflect.Type
	actionValue  reflect.Value
	actionType  reflect.Type
}

func NewTypeDefinition[T Action](def any) (*ActionTypeDefinition, error) {
	if strc, ok := def.(T); ok {
		return TypeDefinitionFromStruct(strc), nil
	}

	if ctor, ok := def.(Constructor[T]); ok {
		return TypeDefinitionFromConstructor(ctor), nil
	}

	return nil, fmt.Errorf("error constructing Action Type Definition")
}

func TypeDefinitionFromConstructor[T Action](defCtor Constructor[T]) *ActionTypeDefinition {
	vCtor := reflect.ValueOf(defCtor)
	tCtor := vCtor.Type()
	
	tAction := tCtor.Out(0)
	tAction = utils.GetValueType(tAction)
	vAction := reflect.New(tAction)

	return &ActionTypeDefinition{
		ctorValue:    vCtor,
		ctorType:    tCtor,
		actionValue:  vAction,
		actionType:  tAction,
	}
}

func TypeDefinitionFromStruct[T Action](def T) *ActionTypeDefinition {
	var ctor Constructor[T] = func() *T { return &def }
	
	vAction := reflect.ValueOf(&def)
	tAction := reflect.TypeOf(def)
	vCtor := reflect.ValueOf(ctor)
	tCtor := vCtor.Type()

	return &ActionTypeDefinition{
		ctorValue: vCtor,
		ctorType: tCtor,
		actionValue: vAction,
		actionType: tAction,
	}
}

