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

func NewAction[T action.GoAction](typedef *ActionTypeDefinition, inst *action.ActionInstance, props *action.GoActionProps) (T, error) {
	act, ok := typedef.ActionValue.Interface().(T)
	if !ok {
		return act, fmt.Errorf("new action does not match expected type")
	}

	act.Init(inst)
	return act, nil
}
