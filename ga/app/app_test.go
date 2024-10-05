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

type expectation struct {
	action *action.GoAction[myAction]
	throws bool
}

func TestNewAction(t *testing.T) {
	app := NewApp()
	def := action.ActionDefinition{}
	app.RegisterActionDef(&def)
	expected := action.NewAction[myAction](&def)

	tests := []cr.TestCase[reflect.Type, expectation]{
		{Name: "valid def type", Input: def.ActionType(), Expected: expectation{action: expected, throws: false}},
		{Name: "invalid def type", Input: reflect.TypeOf(""), Expected: expectation{action: nil, throws: true}},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, expectation]) {
		result, err := NewAction[myAction](test.Input)(app)
		if test.Expected.throws && err == nil {
			t.Errorf("expected error got nil")
			return
		}

		if !test.Expected.throws && *result != *expected {
			t.Errorf("error instatiating action: expected %v, got %v", *expected, *result)
		}
	})
}
