package registration

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	"testing"
)

type myAction struct{}
func (ma myAction) Execute() {}
func myActionCtor(*action.ActionInstance) *myAction {
	return &myAction{}
}

func TestNewRegisteredAction(t *testing.T) {
	registration := &action.GoActionRegistration[myAction]{Constructor: myActionCtor}

	expectedActionDefinition,_ := definition.NewActionDefinition(registration.Constructor)

	registeredAction, err := NewRegisteredAction(registration)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, registration, registeredAction.registration)
	asserts.Equals(t, expectedActionDefinition, registeredAction.ActionDefinition)
}