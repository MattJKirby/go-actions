package action

import (
	"go-actions/ga/utils"
	"reflect"
)

type ActionDefinition struct {
	name     string
	typeName string
	vCtor    reflect.Value
	tCtor    reflect.Type
	vAction  reflect.Value
	tAction  reflect.Type
}

func NewActionDefinition(actionDefCtor any) *ActionDefinition {
	vCtor := reflect.ValueOf(actionDefCtor)
	tCtor := vCtor.Type()

	if tCtor.Kind() != reflect.Func {
		panic("definition constructor must be a function")
	}

	tAction := tCtor.Out(0)
	tAction = utils.GetValueType(tAction)
	vAction := reflect.New(tAction)

	// val, ok := vCtor.Interface().(ActionConstructor[tAction])
	// if ok {
	// 	println(true, val)
	// }

	// val, ok := vAction.Interface().(ActionFunction)
	// if ok {
	// 	println("OKAY", val)
	// }

	// v = v.Call([]reflect.Value{})[0]
	return &ActionDefinition{
		name:     utils.TypeName(tAction),
		typeName: utils.TypePath(tAction),
		vCtor:    vCtor,
		tCtor:    tCtor,
		vAction:  vAction,
		tAction:  tAction,
	}
}

func (ad *ActionDefinition) ActionType() reflect.Type {
	return ad.tAction
}

func (ad *ActionDefinition) ActionValue() reflect.Value {
	return ad.vAction
}

func (ad *ActionDefinition) Constructor() reflect.Value {
	return ad.vCtor
}

func (ad *ActionDefinition) Name() string {
	return ad.name
}

func (ad *ActionDefinition) TypeName() string {
	return ad.typeName
}
