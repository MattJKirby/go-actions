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
	Trigger     bool
	ActionValue reflect.Value
	ActionType  reflect.Type
	PropsValue  reflect.Value
	PropsType   reflect.Type
}

func TypeDefinitionFromRegistration[T action.GoAction, Props action.GoActionProps](reg *action.GoActionRegistration[T, Props]) *ActionTypeDefinition {
	vAction := reflect.ValueOf(reg.Action)
	tAction := vAction.Type()

	tProps := reflect.TypeOf(reg.DefaultProps)
	vProps := reflect.ValueOf(reg.DefaultProps)

	_, Trigger := any(new(T)).(action.TriggerAction)

	return &ActionTypeDefinition{
		TypeName:    utils.TypeName(tAction),
		TypePath:    utils.TypePath(tAction),
		Trigger:     Trigger,
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

func (atd *ActionTypeDefinition) NewAction() (action.GoAction, error) {
	if act, ok := atd.ActionValue.Interface().(action.GoAction); ok {
		return act, nil
	}

	return nil, fmt.Errorf("new action does not match expected type")
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
