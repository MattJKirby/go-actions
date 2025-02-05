package executable

import (
	"go-actions/ga/action/definition"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestNewExecutableAction(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()
	def := definition.NewActionDefinition(&reg)
	typeDef := def.GetTypeDefinition()

	action := NewExecutableAction(typeDef, nil)

	if action == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}