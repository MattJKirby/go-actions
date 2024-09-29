package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/utils"
	"reflect"
)

var ga = app.NewApp()

func DefineAction[T action.Action](actionConstructor action.Constructor[T]) *action.ActionDefinition {
	def := action.NewActionDefinition(actionConstructor)
	return ga.RegisterActionDef(def)
}

func GetActionDefinition(action any) (*action.ActionDefinition, error) {
	actionType := utils.GetValueType(reflect.TypeOf(action))
	return ga.GetActionDef(actionType)
}

func GetAction[T action.Action](a T) (*action.GoAction[T]){
	actionType := utils.GetValueType(reflect.TypeOf(a))
	action, err := app.NewAction[T](actionType)(ga)
	if err != nil {
		panic("aaa")
	}

	return action
}
