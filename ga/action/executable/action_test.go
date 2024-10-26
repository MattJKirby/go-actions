package executable

import (
	"go-actions/ga/action"
	"testing"
)

type testAction struct {}
func (ta testAction) Execute(){}

func TestNewAction(t *testing.T){
	def := &action.ActionDefinition{}
	acn := NewAction[testAction](def)
	t.Run("test new action", func(t *testing.T) {
		if acn.Instance == nil {
			t.Errorf("invalid action instance: expected %v got %v", action.NewActionInstance(def), acn)
		}
	})
}