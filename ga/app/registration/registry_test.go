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
	registration := ta.GenerateActionValidEmptyRegistration()

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

func TestGetActionByType(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidEmptyRegistration()
	def, _ := definition.NewActionDefinition(&registration)

	AcceptRegistration(&registration)(registry)

	tests := []cr.TestCase[reflect.Type, *definition.ActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps]]{
		{Name: "existing def", Input: def.ActionType, Expected: def},
		{Name: "not existing def", Input: reflect.TypeOf("err"), Expected: nil, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, *definition.ActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps]]) {
		storedDef, err := GetTypedActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps](test.Input)(registry)
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

func TestGetActionByName(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidRegistration()
	AcceptRegistration(&registration)(registry)

	def, _ := definition.NewActionDefinition(&registration)

	tests := []struct {
		name     string
		input    string
		expected *definition.ActionTypeDefinition
		hasError bool
	}{
		{name: "existing action", input: "ActionValid", expected: def.ActionTypeDefinition, hasError: false},
		{name: "non existing action", input: "xxxx", expected: nil, hasError: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			result, err := GetRegisteredTypeDefinitionByName(test.input)(registry)
			if test.expected != nil {
				asserts.Equals(t, test.expected.ActionType, result.ActionType)
			}

			asserts.Equals(t, test.hasError, err != nil)
		})
	}

}
