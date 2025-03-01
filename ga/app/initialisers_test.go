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
	definition := definition.TypeDefinitionFromRegistration[ta.ActionValidEmpty, ta.ActionValidEmptyProps](&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockConfig)
	expected := &InitialisedAction{
		Action:   reg.Constructor(instance, ta.ActionValidEmptyProps{}),
		InitialisedInstance: instance,
	}

	actual, err := InitialiseNewAction(mockConfig, definition)
	asserts.Equals(t, expected, actual)
	asserts.Equals(t, nil, err)
}

// func TestInitialiseTypedAction(t *testing.T) {
// 	app, reg := appWithEmptyRegistration(mockConfig)
// 	def := definition.NewActionDefinition(&reg)

// 	instance := action.NewActionInstance(def.TypeName, mockConfig)
// 	action := reg.Constructor(instance, *reg.DefaultProps)
// 	expectedInstantiatedTypedAction := &InstantiatedTypedAction[ta.ActionValidEmpty]{
// 		Instance: instance,
// 		Action:   action,
// 	}

// 	actual, err := InitialiseTypedAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)

// 	asserts.Equals(t, expectedInstantiatedTypedAction, actual)
// 	asserts.Equals(t, nil, err)
// }