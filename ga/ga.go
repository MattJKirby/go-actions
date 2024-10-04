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

func GetActionDefinition[T action.Action](action T) (*action.ActionDefinition, error) {
	actionType := utils.GetValueType(reflect.TypeOf(action))
	return ga.GetActionDef(actionType)
}

func GetAction[T action.Action](a T) (*action.GoAction[T], error){
	actionType := utils.GetValueType(reflect.TypeOf(a))
	return app.NewAction[T](actionType)(ga)
}

// func GetActionByName(name string)
