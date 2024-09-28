package ga

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/utils"
	"reflect"
)

var ga = app.NewApp()

func DefineAction[T action.FunctionDefinition](actionConstructor action.Constructor[T]) *action.ActionDefinition {
	def := action.NewActionDefinition(actionConstructor)
	return ga.RegisterActionDef(def)
}

func GetAction(action any) (*action.ActionDefinition, error) {
	actionType := utils.GetValueType(reflect.TypeOf(action))
	return ga.GetActionDef(actionType)
}
