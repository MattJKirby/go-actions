package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestInitialiseNewAction(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()
	definition := definition.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockConfig)

	expected := &InitialisedAction{
		Action:              reg.Constructor(instance, ta.ActionValidEmptyProps{}),
		InitialisedInstance: instance,
	}

	actual, err := InitialiseNewAction(mockConfig, definition)
	asserts.Equals(t, expected, actual)
	asserts.Equals(t, nil, err)
}

func TestInitialiseTypedAction(t *testing.T) {
	app, reg := appWithEmptyRegistration(mockConfig)
	def := definition.NewActionDefinition(&reg)

	instance := action.NewActionInstance(def.TypeName, mockConfig)
	expectedInstantiatedTypedAction := &InitialisedTypedAction[ta.ActionValidEmpty]{
		Action:              reg.Constructor(instance, *reg.DefaultProps),
		InitialisedInstance: instance,
	}

	actual, err := InitialiseNewTypedAction(app.modelConfig, def)

	asserts.Equals(t, expectedInstantiatedTypedAction, actual)
	asserts.Equals(t, nil, err)
}
