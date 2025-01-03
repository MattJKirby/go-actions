package registration

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewRegisteredAction(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid, th.ActionValidProps]()
	registration := &action.GoActionRegistration[th.ActionValid, th.ActionValidProps]{Constructor: ctor}

	expectedActionDefinition, _ := definition.NewActionDefinition(registration.Constructor)

	registeredAction, err := NewRegisteredAction(registration)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, registration, registeredAction.Registration)
	asserts.Equals(t, expectedActionDefinition, registeredAction.ActionDefinition)
}
