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

func defHelper() *definition.ActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps] {
	reg := ta.GenerateActionValidEmptyRegistration()
	return definition.NewActionDefinition(&reg)
}

func TestNewExecutableAction(t *testing.T) {
	def := defHelper()
	typeDef := def.GetTypeDefinition()

	expectedInst := action.NewActionInstance(def.TypeName, mockConfig)
	def.Constructor(expectedInst, ta.ActionValidEmptyProps{})

	executableAction := NewExecutableAction(mockConfig, typeDef, nil)

	asserts.Equals(t, expectedInst, executableAction.instance)
}

func TestNewExecutableInstance(t *testing.T) {
	def := defHelper()
	typeDef := def.GetTypeDefinition()

	expectedInst := action.NewActionInstance(def.TypeName, mockConfig)
	def.Constructor(expectedInst, ta.ActionValidEmptyProps{})

	inst := newExecutableInstance(mockConfig, typeDef)

	asserts.Equals(t, expectedInst, inst)
}
