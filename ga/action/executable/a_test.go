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

func defHelper() *definition.ActionDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps]{
	reg := ta.GenerateActionValidEmptyRegistration()
	return definition.NewActionDefinition(&reg)
}

func TestNewExecutableAction(t *testing.T) {
	def := defHelper()
	typeDef := def.GetTypeDefinition()
	action := NewExecutableAction(mockConfig, typeDef, nil)

	if action == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}

func TestNewExeutableInstance(t *testing.T) {
	def := defHelper()
	typeDef := def.GetTypeDefinition()

	expectedInst := action.NewActionInstance(def.TypeName, mockConfig)
	def.Constructor(expectedInst, ta.ActionValidEmptyProps{})
	
	inst := newExecutableInstance(mockConfig, typeDef)


	asserts.Equals(t, expectedInst, inst)
}