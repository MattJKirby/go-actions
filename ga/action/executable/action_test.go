package executable

import (
	"go-actions/ga/action/definition"
	"go-actions/ga/action/instance"
	"testing"
)

type testAction struct {}
func (ta testAction) Execute(){}

func TestNewAction(t *testing.T){
	def := &definition.ActionDefinition{}
	acn := NewAction[testAction](def)
	t.Run("test new action", func(t *testing.T) {
		if acn.Instance == nil {
			t.Errorf("invalid action instance: expected %v got %v", instance.NewActionInstance(def.Name), acn)
		}
	})
}