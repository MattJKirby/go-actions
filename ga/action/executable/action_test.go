package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewAction(t *testing.T) {

	reg := &action.GoActionRegistration[th.ActionValid, th.ActionValidProps]{}
	def, _ := definition.NewActionDefinition[th.ActionValid, th.ActionValidProps](reg)

	acn := NewAction(*def)

	if acn.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}
