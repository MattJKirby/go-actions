package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"reflect"
)

var ga = app.NewApp()

func DefineAction[T action.FunctionDefinition](actionConstructor action.Constructor[T]) *action.ActionDefinition {
	return ga.RegisterActionDef(action.NewActionDefinition(actionConstructor))
}

func GetAction(def any) (*action.ActionDefinition, error) {
	v := reflect.ValueOf(def)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return ga.GetActionDef(v)
}
