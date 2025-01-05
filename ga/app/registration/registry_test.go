package registration

import (
	"go-actions/ga/action/definition"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"reflect"
	"testing"
)

func TestAcceptAction(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidRegistration()

	err := AcceptRegistration(&registration)(registry)
	abt := len(registry.actionsByType)
	abn := len(registry.actionsByName)

	if abt != 1 {
		t.Errorf("test actions by type: got: %d, expected: %d", abt, 1)
	}

	if abn != 1 {
		t.Errorf("test actions by name: got: %d, expected: %d", abt, 1)
	}

	asserts.Equals(t, nil, err)
}

func TestGetAction(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidRegistration()
	def, _ := definition.NewActionDefinition(&registration)

	AcceptRegistration(&registration)(registry)

	tests := []cr.TestCase[reflect.Type, *definition.ActionDefinition[ta.ActionValid, ta.ActionValidProps]]{
		{Name: "existing def", Input: def.ActionType, Expected: def},
		{Name: "not existing def", Input: reflect.TypeOf("err"), Expected: nil, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, *definition.ActionDefinition[ta.ActionValid, ta.ActionValidProps]]) {
		storedDef, err := GetAction[ta.ActionValid, ta.ActionValidProps](test.Input)(registry)
		hasErr := err != nil

		if test.Error != hasErr {
			t.Errorf("test %s: expected an error but got none", test.Name)
			return
		}

		if !test.Error && storedDef.Name != test.Expected.Name {
			t.Errorf("test %s: got %v, expected %v", test.Name, storedDef, test.Expected)
		}
	})
}
