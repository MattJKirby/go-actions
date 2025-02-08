package definition

import (
	"fmt"
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

type ActionConstructor func(*action.ActionInstance, action.GoActionProps) (action.GoAction, error)

func TypeDefinitionFromRegistration[T action.GoAction, Props action.GoActionProps](reg *action.GoActionRegistration[T, Props]) *ActionTypeDefinition {
	vCtor := reflect.ValueOf(reg.Constructor)
	tCtor := vCtor.Type()

	tProps := reflect.TypeOf(*reg.DefaultProps)
	vProps := reflect.ValueOf(*reg.DefaultProps)

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

func (atd *ActionTypeDefinition) NewDefaultProps() action.GoActionProps {
	return atd.PropsValue.Interface()
}

func (atd *ActionTypeDefinition) ValidatePropsType(props action.GoActionProps) error {
	propsType := reflect.TypeOf(props)

	switch {
	case propsType == nil:
		return fmt.Errorf("props can't be nil")

	case propsType.Kind() == reflect.Pointer:
		return fmt.Errorf("props must be value type")

	case propsType != atd.PropsType:
		return fmt.Errorf("props type does not match registered default props type")
	}

	return nil
}

func (atd *ActionTypeDefinition) NewConstructor() ActionConstructor {
	return func(instance *action.ActionInstance, props action.GoActionProps) (action.GoAction, error) {
		if err := atd.ValidatePropsType(props); err != nil {
			return nil, err
		}

		results := atd.CtorValue.Call([]reflect.Value{
			reflect.ValueOf(instance),
			reflect.ValueOf(props),
		})
		return results[0].Interface().(action.GoAction), nil
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
