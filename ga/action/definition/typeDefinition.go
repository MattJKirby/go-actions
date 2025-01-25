package definition

import (
	"go-actions/ga/action"

	"go-actions/ga/utils"
	"reflect"
)

type ActionTypeDefinition struct {
	CtorValue   reflect.Value
	CtorType    reflect.Type
	ActionValue reflect.Value
	ActionType  reflect.Type
	PropsValue  reflect.Value
	PropsType   reflect.Type
}

func TypeDefinitionFromRegistration[T action.GoAction, Props action.GoActionProps](reg *action.GoActionRegistration[T, Props]) *ActionTypeDefinition {
	vCtor := reflect.ValueOf(reg.Constructor)
	tCtor := vCtor.Type()

	vProps := reflect.ValueOf(reg.DefaultProps)
	tProps := vProps.Type()

	tAction := tCtor.Out(0)
	tAction = utils.GetValueType(tAction)
	vAction := reflect.New(tAction)

	return &ActionTypeDefinition{
		CtorValue:   vCtor,
		CtorType:    tCtor,
		ActionValue: vAction,
		ActionType:  tAction,
		PropsValue:  vProps,
		PropsType:   tProps,
	}
}

// func TypeDefinitionFromStruct[T action.GoAction, Props action.GoActionProps](def T) *ActionTypeDefinition {
// 	var ctor action.GoActionConstructor[T, Props] = func(*action.ActionInstance, Props) *T { return &def }

// 	vAction := reflect.ValueOf(&def)
// 	tAction := reflect.TypeOf(def)
// 	vCtor := reflect.ValueOf(ctor)
// 	tCtor := vCtor.Type()

// 	return &ActionTypeDefinition{
// 		CtorValue:   vCtor,
// 		CtorType:    tCtor,
// 		ActionValue: vAction,
// 		ActionType:  tAction,
// 	}
// }
