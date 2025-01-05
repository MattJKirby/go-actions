package executable

import (
	"go-actions/ga/action/definition"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestNewAction(t *testing.T) {

	reg := ta.GenerateActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)

	acn := NewAction(*def)

	if acn.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}
