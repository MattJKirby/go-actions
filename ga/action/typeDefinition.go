package action

import (
	"go-actions/ga/utils"
	"reflect"
)

type TypeDefinition struct {
	TypeName    string
	TypePath    string
	Trigger     bool
	ActionValue reflect.Value
	ActionType  reflect.Type
}

func TypeDefinitionFromRegistration[T GoAction](reg *ActionRegistration[T]) *TypeDefinition {
	vAction := reflect.ValueOf(reg.Action)
	tAction := vAction.Type()

	// tProps := reflect.TypeOf(reg.DefaultProps)
	// vProps := reflect.ValueOf(reg.DefaultProps)

	_, Trigger := any(new(T)).(TriggerAction)

	return &TypeDefinition{
		TypeName:    utils.TypeName(tAction),
		TypePath:    utils.TypePath(tAction),
		Trigger:     Trigger,
		ActionValue: vAction,
		ActionType:  tAction,
		// PropsValue:  vProps,
		// PropsType:   tProps,
	}
}

// func (atd *ActionTypeDefinition) NewDefaultProps() action.GoActionProps {
// 	return atd.PropsValue.Interface()
// }

// func (atd *ActionTypeDefinition) ValidatePropsType(props action.GoActionProps) error {
// 	propsType := reflect.TypeOf(props)

// 	switch {
// 	case propsType == nil:
// 		return fmt.Errorf("props can't be nil")

// 	case propsType.Kind() == reflect.Pointer:
// 		return fmt.Errorf("props must be value type")

// 	case propsType != atd.PropsType:
// 		return fmt.Errorf("props type does not match registered default props type")
// 	}

// 	return nil
// }

// func InitialiseInstance[T GoAction](typedef *TypeDefinition, inst *ActionInstance) (T, error) {

// 	return act, nil
// }
