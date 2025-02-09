package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"go-actions/ga/testing/testHelpers/actionModelTestHelpers"
	"testing"
)

var mockConfig = &actionModelTestHelpers.MockActionModelConfig{MockUid: "uid"}

func defHelper() *definition.ActionDefinition[ta.ActionValid, ta.ActionValidProps] {
	reg := ta.GenerateActionValidRegistration()
	return definition.NewActionDefinition(&reg)
}

func TestNewExecutableAction(t *testing.T) {
	def := defHelper()
	typeDef := def.GetTypeDefinition()

	expectedInst := action.NewActionInstance(def.TypeName, mockConfig)
	def.Constructor(expectedInst, ta.ActionValidDefaultProps)

	executableAction := NewExecutableAction(mockConfig, typeDef)

	asserts.Equals(t, expectedInst, executableAction.instance)
}

func TestNewExecutableInstance(t *testing.T) {

	def := defHelper()
	typeDef := def.GetTypeDefinition()
	expectedInst := action.NewActionInstance(def.TypeName, mockConfig)
	def.Constructor(expectedInst, ta.ActionValidDefaultProps)

	tests := []struct {
		name             string
		inputProps       any
		expectedInstance *action.ActionInstance
		expectErr        bool
	}{
		{name: "with valid empty props", inputProps: ta.ActionValidDefaultProps, expectedInstance: expectedInst, expectErr: false},
		{name: "with valid props", inputProps: ta.ActionValidDefaultProps, expectedInstance: expectedInst, expectErr: false},
		{name: "with invalid props", inputProps: ta.ActionInvalidNoExecute{}, expectedInstance: nil, expectErr: true},
		{name: "with nil props", inputProps: nil, expectedInstance: expectedInst, expectErr: false},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			inst, _, err := newExecutableInstance(mockConfig, typeDef, test.inputProps)
			hasErr := err != nil

			asserts.Equals(t, test.expectErr, hasErr)
			asserts.Equals(t, test.expectedInstance, inst)
		})
	}
}
