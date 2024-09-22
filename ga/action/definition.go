package action

import (
	"go-actions/ga/types"
	"reflect"
	"strings"
)

type ActionDefinition struct {
	name string
	typeName string
	v reflect.Value
	t reflect.Type
}

func NewActionDefinition(defObj any) *ActionDefinition {
	v := reflect.ValueOf(defObj)
	t := v.Type()

	if !types.IsRefType(t) {
		panic("definition must be a ref type")
	}

	if t.Kind() == reflect.Func {
		// v = v.Call([]reflect.Value{})[0]
		t = t.Out(0)

		if types.IsRefType(t) {
			t = t.Elem()
		}
	}

	s := strings.Split(t.String(), ".")
	name := s[len(s) -1]


	return &ActionDefinition{
		name:     name,
		typeName: types.TypeName(t),
		v:        v,
		t:        t,
	}
}


func (ad *ActionDefinition) Type() reflect.Type {
	return ad.t
}

func (ad *ActionDefinition) Value() reflect.Value {
	return ad.v
}

func (ad *ActionDefinition) Name() string {
	return ad.name
}

func (ad *ActionDefinition) TypeName() string {
	return ad.typeName
}


