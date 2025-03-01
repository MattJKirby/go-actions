package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestNewInitialisedAction(t *testing.T) {
	app, reg := appWithEmptyRegistration(mockConfig)
	definition := definition.TypeDefinitionFromRegistration[ta.ActionValidEmpty, ta.ActionValidEmptyProps](&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockConfig)
	expected := &InitialisedAction{
		Action:   reg.Constructor(instance, ta.ActionValidEmptyProps{}),
		InitialisedInstance: instance,
	}

	actual, err := NewInitialisedAction(app, definition)
	asserts.Equals(t, expected, actual)
	asserts.Equals(t, nil, err)
}