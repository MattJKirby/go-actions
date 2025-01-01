package registration

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewRegisteredAction(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid]()
	registration := &action.GoActionRegistration[th.ActionValid]{Constructor: ctor}

	expectedActionDefinition,_ := definition.NewActionDefinition(registration.Constructor)

	registeredAction, err := NewRegisteredAction(registration)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, registration, registeredAction.registration)
	asserts.Equals(t, expectedActionDefinition, registeredAction.ActionDefinition)
}