package definition

import (
	"go-actions/ga/action"

	"go-actions/ga/utils"
	"reflect"
)

type ActionTypeDefinition struct {
	TypeName    string
	TypePath    string
	CtorValue   reflect.Value
	CtorType    reflect.Type
	ActionValue reflect.Value
	ActionType  reflect.Type
	PropsValue  reflect.Value
	PropsType   reflect.Type
}

type ActionConstructor func(*action.ActionInstance, action.GoActionProps) action.GoAction

func TypeDefinitionFromRegistration[T action.GoAction, Props action.GoActionProps](reg *action.GoActionRegistration[T, Props]) *ActionTypeDefinition {
	vCtor := reflect.ValueOf(reg.Constructor)
	tCtor := vCtor.Type()

	vProps := reflect.ValueOf(reg.DefaultProps)
	tProps := vProps.Type()

	tAction := tCtor.Out(0)
	tAction = utils.GetValueType(tAction)
	vAction := reflect.New(tAction)

	return &ActionTypeDefinition{
		TypeName:    utils.TypeName(tAction),
		TypePath:    utils.TypePath(tAction),
		CtorValue:   vCtor,
		CtorType:    tCtor,
		ActionValue: vAction,
		ActionType:  tAction,
		PropsValue:  vProps,
		PropsType:   tProps,
	}
}

func (atd *ActionTypeDefinition) NewDefaultProps() any {
	return atd.PropsValue.Interface()
}

func (atd *ActionTypeDefinition) NewConstructor() ActionConstructor {
	callable := func(instance *action.ActionInstance, props action.GoActionProps) action.GoAction {
		results := atd.CtorValue.Call([]reflect.Value{
			reflect.ValueOf(instance),
			reflect.ValueOf(props),
		})
		return results[0].Interface().(action.GoAction)
	}

	return callable
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
