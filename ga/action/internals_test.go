package action_test

import (
	"go-actions/ga/action"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"testing"
)

func TestNewActionInternals(t *testing.T) {
	reg := action.ActionRegistration[testActions.ActionValidEmpty]{Action: testActions.ActionValidEmpty{}}
	definition := action.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance(mockConfig, definition)
	
	internals := action.NewInternals(instance)

	assert.Equals(t, instance, internals.GetInstance(), )
	assert.Equals(t, true, internals.GetInput() != nil)
	assert.Equals(t, true, internals.GetOutput() != nil)
}

func TestInitInternals(t *testing.T) {
	internals := &action.Internals{}
	internals.SetInstance(&action.ActionInstance{Name: "a"})
	newInternals := &action.Internals{}
	newInternals.SetInstance(&action.ActionInstance{Name: "b"})

	internals.InitInternals(newInternals)

	assert.Equals(t, internals, newInternals)

}