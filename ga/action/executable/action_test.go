package executable

import (
	"go-actions/ga/action/definition"
	"testing"
)

type testAction struct{}
type testActionProps struct{}

func (ta testAction) Execute() {}

func TestNewAction(t *testing.T) {
	def := &definition.ActionDefinition{}
	acn := NewAction[testAction, testActionProps](def)

	if acn.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}
