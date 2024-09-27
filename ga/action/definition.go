package action

import (
	"fmt"
	"go-actions/ga/types"
	"reflect"
	"strings"
)

type ActionDefinition struct {
	name     string
	typeName string
	vCtor    reflect.Value
	tCtor    reflect.Type
	tAction reflect.Type
}

func NewActionDefinition(actionDefCtor any) *ActionDefinition {
	vCtor := reflect.ValueOf(actionDefCtor)
	tCtor := vCtor.Type()

	if tCtor.Kind() != reflect.Func {
		panic("definition constructor must be a function")
	}

	tAction := tCtor.Out(0)
	if types.IsRefType(tAction) {
		tAction = tAction.Elem()
	}

	s := strings.Split(tCtor.String(), ".")
	name := s[len(s)-1]

	fmt.Println(name, types.TypeName(tAction))

	// v = v.Call([]reflect.Value{})[0]
	return &ActionDefinition{
		name:     name,
		typeName: types.TypeName(tAction),
		vCtor:    vCtor,
		tCtor:    tCtor,
		tAction:  tAction,
	}
}

func (ad *ActionDefinition) ActionType() reflect.Type {
	return ad.tAction
}

func (ad *ActionDefinition) ActionConstructor() reflect.Value {
	return ad.vCtor
}

func (ad *ActionDefinition) Name() string {
	return ad.name
}

func (ad *ActionDefinition) TypeName() string {
	return ad.typeName
}
