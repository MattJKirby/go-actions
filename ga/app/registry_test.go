package app

import (
	"go-actions/ga/action"
	"go-actions/ga/cr"
	"reflect"
	"testing"
)

func TestAcceptDefinition(t *testing.T) {
	defReg := NewActionDefinitionRegistry()
	def := action.ActionDefinition{}

	defReg.acceptDefinition(&def)
	abt := len(defReg.actionsByType)
	abn := len(defReg.actionsByName)

	if abt != 1 {
		t.Errorf("test actions by type: got: %d, expected: %d", abt, 1)
	}

	if abn != 1 {
		t.Errorf("test actions by name: got: %d, expected: %d", abt, 1)
	}
}

func TestGetDefinition(t *testing.T) {
	defReg := NewActionDefinitionRegistry()
	def := action.ActionDefinition{}
	defReg.acceptDefinition(&def)

	tests := []cr.TestCase[reflect.Type, *action.ActionDefinition]{
		{Name: "existing def", Input: def.ActionType(), Expected: &def},
		{Name: "not existing def", Input: reflect.TypeOf("err"), Expected: nil, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, *action.ActionDefinition]) {
		storedDef, err := defReg.getDefinition(test.Input)

		if test.Error && err == nil {
			t.Errorf("test %s: expected an error but got none", test.Name)
			return
		}

		if !test.Error && storedDef != test.Expected {
			t.Errorf("test %s: got %v, expected %v", test.Name, storedDef, test.Expected)
		}
	})
}
