package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"reflect"
)

var ga = app.NewApp()

func DefineAction(actionConstructor any) *action.ActionDefinition {
	return ga.RegisterActionDef(action.NewActionDefinition(actionConstructor))
}

func GetAction(def any) (*action.ActionDefinition, error) {
	v := reflect.ValueOf(def)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	
	return ga.GetAction(v)
}