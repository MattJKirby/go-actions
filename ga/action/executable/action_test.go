package executable

import (
	"go-actions/ga/action/definition"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewAction(t *testing.T) {
	def := &definition.ActionDefinition{}
	acn := NewAction[th.ActionValid](def)

	if acn.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}
