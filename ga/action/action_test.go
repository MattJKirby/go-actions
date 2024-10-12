package action

import (
	"testing"
)

type testAction struct {}
func (ta testAction) Execute(){}

func TestNewAction(t *testing.T){
	def := ActionDefinition{}
	inst := ActionInstance{}
	acn := NewAction[testAction](&def, &inst)

	t.Run("test new action", func(t *testing.T) {
		if acn.Instance != &inst {
			t.Errorf("invalid action instance: expected %v got %v", inst, acn.Instance)
		}
	})
}