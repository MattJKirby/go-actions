package app

import (
	"go-actions/ga/action"
	"testing"
)

func TestRegisterDef(t *testing.T){
	app := NewApp()
	def := action.ActionDefinition{}

	result := app.RegisterActionDef(&def)

	t.Run("Test register def", func(t *testing.T) {
		if result != &def {
			t.Errorf("Error during registration: expected %v, got %v", &def, result)
		}
	})
}

func TestGetActionDef(t *testing.T){
	app := NewApp()
	def := action.ActionDefinition{}
	app.RegisterActionDef(&def)

	result, _ := app.GetActionDef(def.ActionType())

	t.Run("Test get def", func(t *testing.T) {
		if result != &def {
			t.Errorf("Error getting definition: expected %v, got %v", &def, result)
		}
	})
}

type myAction struct {}
func (ma myAction) Execute(){}

func TestNewAction(t *testing.T){
	app := NewApp()
	def := action.ActionDefinition{}
	app.RegisterActionDef(&def)

	expected := action.NewAction[myAction](&def)
	result, err := NewAction[myAction](def.ActionType())(app)

	t.Run("Test making action", func(t *testing.T) {
		if *result != *expected {
			t.Errorf("error instatiating action: expected %v, got %v", *expected, *result)
		}

		if err != nil {
			t.Errorf("error instatiating action: err not expected %v", err)
		}
	})
}