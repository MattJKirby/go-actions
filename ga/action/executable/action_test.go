package executable

import (
	"go-actions/ga/action/definition"
	"go-actions/ga/app/registration"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewAction(t *testing.T) {
	actionDefiniton := &definition.ActionDefinition{}
	RegisteredAction := registration.RegisteredAction[th.ActionValid, th.ActionValidProps]{ActionDefinition: actionDefiniton}
	acn := NewAction(RegisteredAction)

	if acn.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}
