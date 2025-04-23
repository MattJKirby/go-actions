package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
	"testing"
)

func TestInitialiseNewAction(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()
	definition := definition.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockGlobalConfig)

	expected := &InitialisedAction{
		Action:              reg.Constructor(instance, ta.ActionValidEmptyProps{}),
		InitialisedInstance: instance,
	}

	actual, err := InitialiseNewAction(mockGlobalConfig, definition)
	assert.Equals(t, expected, actual)
	assert.Equals(t, nil, err)
}
