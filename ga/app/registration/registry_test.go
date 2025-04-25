package registration

import (
	"fmt"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
	"testing"
)

func TestAcceptAction(t *testing.T) {
	registry := NewActionRegistry()
	registration := ta.GenerateActionValidEmptyRegistration()

AcceptRegistration(&registration)(registry)
	abt := len(registry.actionsByType)
	abn := len(registry.actionsByName)

	if abt != 1 {
		t.Errorf("test actions by type: got: %d, expected: %d", abt, 1)
	}

	if abn != 1 {
		t.Errorf("test actions by name: got: %d, expected: %d", abt, 1)
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
			result, err := GetTypeDefinitionByName(test.input)(registry)
			fmt.Println(err)
			assert.Equals(t, test.returnsNil, result == nil)
			assert.Equals(t, test.hasError, err != nil)
		})
	}
}
