package app

import (
	"go-actions/ga/action"
	"go-actions/ga/cr"
	"reflect"
	"testing"
)

func TestRegisterDef(t *testing.T) {
	app := NewApp()
	def := action.ActionDefinition{}

	result := app.RegisterActionDef(&def)

	t.Run("Test register def", func(t *testing.T) {
		if result != &def {
			t.Errorf("Error during registration: expected %v, got %v", &def, result)
		}
	})
}

func TestGetActionDef(t *testing.T) {
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

type myAction struct{}

func (ma myAction) Execute() {}

func TestNewAction(t *testing.T) {
	app := NewApp()
	def := action.ActionDefinition{}
	app.RegisterActionDef(&def)

	tests := []cr.TestCase[reflect.Type, *action.GoAction[myAction]]{
		{Name: "valid def type", Input: def.ActionType()},
		{Name: "invalid def type", Input: reflect.TypeOf(""), Expected: nil, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, *action.GoAction[myAction]]) {
		result, err := NewAction[myAction](test.Input)(app)
		if test.Error && err == nil {
			t.Errorf("expected error got nil")
			return
		}

		if !test.Error && result == nil {
			t.Errorf("error instatiating action: got %v", nil)
		}
	})
}
