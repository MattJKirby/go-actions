package action

import (
	"fmt"
	"go-actions/ga/utils"
	"reflect"
)

type ActionTypeDefinition struct {
	CtorValue    reflect.Value
	CtorType    reflect.Type
	ActionValue  reflect.Value
	ActionType  reflect.Type
}

func NewTypeDefinition[T GoAction](def any) (*ActionTypeDefinition, error) {
	if strc, ok := def.(T); ok {
		return TypeDefinitionFromStruct(strc), nil
	}

	if ctor, ok := def.(GoActionConstructor[T]); ok {
		return TypeDefinitionFromConstructor(ctor), nil
	}

	return nil, fmt.Errorf("error constructing Action Type Definition")
}

func TypeDefinitionFromConstructor[T GoAction](defCtor GoActionConstructor[T]) *ActionTypeDefinition {
	vCtor := reflect.ValueOf(defCtor)
	tCtor := vCtor.Type()
	
	tAction := tCtor.Out(0)
	tAction = utils.GetValueType(tAction)
	vAction := reflect.New(tAction)

	return &ActionTypeDefinition{
		CtorValue:    vCtor,
		CtorType:    tCtor,
		ActionValue:  vAction,
		ActionType:  tAction,
	}
}

func TypeDefinitionFromStruct[T GoAction](def T) *ActionTypeDefinition {
	var ctor GoActionConstructor[T] = func() *T { return &def }
	
	vAction := reflect.ValueOf(&def)
	tAction := reflect.TypeOf(def)
	vCtor := reflect.ValueOf(ctor)
	tCtor := vCtor.Type()

	return &ActionTypeDefinition{
		CtorValue: vCtor,
		CtorType: tCtor,
		ActionValue: vAction,
		ActionType: tAction,
	}
}

