package definition

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/utils"
	"reflect"
)

type ActionTypeDefinition struct {
	CtorValue    reflect.Value
	CtorType    reflect.Type
	ActionValue  reflect.Value
	ActionType  reflect.Type
}

func NewTypeDefinition[T action.GoAction](def any) (*ActionTypeDefinition, error) {
	if strc, ok := def.(T); ok {
		return TypeDefinitionFromStruct(strc), nil
	}

	if ctor, ok := def.(action.GoActionConstructor[T]); ok {
		return TypeDefinitionFromConstructor(ctor), nil
	}

	return nil, fmt.Errorf("error constructing Action Type Definition")
}

func TypeDefinitionFromConstructor[T action.GoAction](defCtor action.GoActionConstructor[T]) *ActionTypeDefinition {
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

func TypeDefinitionFromStruct[T action.GoAction](def T) *ActionTypeDefinition {
	var ctor action.GoActionConstructor[T] = func(action.GoActionInternals) *T { return &def }
	
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

