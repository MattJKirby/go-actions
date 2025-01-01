package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewRegisteredAction(t *testing.T) {
	expectedRegistration := &action.GoActionRegistration[myAction]{Constructor: myActionCtor}
	expectedActionDefinition,_ := definition.NewActionDefinition(expectedRegistration.Constructor)

	registeredAction, err := NewRegisteredAction(expectedRegistration)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, expectedRegistration, registeredAction.registration)
	asserts.Equals(t, expectedActionDefinition, registeredAction.actionDefinition)
}