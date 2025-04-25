package registration

import (
	"fmt"
	"go-actions/ga/action/definition"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
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

	assert.Equals(t, nil, err)
}

func TestGetActionByType(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidEmptyRegistration()
	def := definition.NewActionDefinition(&registration)

	AcceptRegistration(&registration)(registry)

	tests := []struct {
		name     string
		input    reflect.Type
		expected *definition.ActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps]
		err      bool
	}{
		{name: "existing def", input: def.ActionType, expected: def},
		{name: "not existing def", input: reflect.TypeOf("err"), expected: nil, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storedDef, err := GetTypedActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps](test.input)(registry)
			hasErr := err != nil

			if test.err != hasErr {
				t.Errorf("test %s: expected an error but got none", test.name)
				return
			}

			if !test.err && storedDef.Name != test.expected.Name {
				t.Errorf("test %s: got %v, expected %v", test.name, storedDef, test.expected)
			}
		})
	}
}

func TestGetActionByName(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidRegistration()
	AcceptRegistration(&registration)(registry)

	tests := []struct {
		name       string
		input      string
		returnsNil bool
		hasError   bool
	}{
		{name: "existing action", input: "ActionValid", returnsNil: false, hasError: false},
		{name: "non existing action", input: "xxxx", returnsNil: true, hasError: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := GetRegisteredTypeDefinitionByName(test.input)(registry)
			fmt.Println(err)
			assert.Equals(t, test.returnsNil, result == nil)
			assert.Equals(t, test.hasError, err != nil)
		})
	}
}
