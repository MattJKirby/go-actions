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
	tests := []struct {
		name          string
		inputProps    any
		expectedProps ta.ActionValidProps
	}{
		{name: "with valid empty props", inputProps: ta.ActionValidProps{}, expectedProps: ta.ActionValidProps{}},
		{name: "with valid props", inputProps: ta.ActionValidProps{Param1: "ASDF"}, expectedProps: ta.ActionValidProps{Param1: "ASDF"}},
		{name: "with invalid props", inputProps: ta.ActionInvalidNoExecute{}, expectedProps: ta.ActionValidDefaultProps},
		{name: "with nil props", inputProps: nil, expectedProps: ta.ActionValidDefaultProps},
	}

	def := defHelper()
	typeDef := def.GetTypeDefinition()

	for _, test := range tests {
		t.Helper()
		t.Run(test.name, func(t *testing.T) {

			expectedInst := action.NewActionInstance(def.TypeName, mockConfig)
			def.Constructor(expectedInst, test.expectedProps)

			inst := newExecutableInstance(mockConfig, typeDef, test.inputProps)

			asserts.Equals(t, expectedInst, inst)
		})
	}
}
