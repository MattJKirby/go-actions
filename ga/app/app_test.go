package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"go-actions/ga/testing/testHelpers/actionModelTestHelpers"

	"testing"
)

var mockConfig = &actionModelTestHelpers.MockActionModelConfig{MockUid: "uid"}

func appWithEmptyRegistration(config model.ActionModelConfig) (*App, action.GoActionRegistration[ta.ActionValidEmpty, ta.ActionValidEmptyProps]) {
	app := NewApp("test")
	app.modelConfig = config
	registration := ta.GenerateActionValidEmptyRegistration()
	RegisterAction(&registration)(app)
	return app, registration
}

func TestRegisterActionAndGet(t *testing.T) {
	app, _ := appWithEmptyRegistration(mockConfig)
	result, _ := GetActionRegistration[ta.ActionValidEmpty, ta.ActionValidEmptyProps]()(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestInstantiateAction(t *testing.T) {
	app, reg := appWithEmptyRegistration(mockConfig)

	expectedInstance := action.NewActionInstance("ActionValidEmpty", mockConfig)
	expectedAction := reg.Constructor(expectedInstance, ta.ActionValidEmptyProps{})
	expected := &InstantiatedAction{
		Instance: expectedInstance,
		Action:   expectedAction,
	}

	tests := []struct {
		name      string
		inputName string
		expected  *InstantiatedAction
		expectErr bool
	}{
		{name: "valid - existing action name", inputName: "ActionValidEmpty", expected: expected, expectErr: false},
		{name: "invalid - not existing action name", inputName: "notregistered", expected: nil, expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := InstantiateActionFromName(test.inputName)(app)
			hasErr := err != nil

			asserts.Equals(t, test.expected, actual)
			asserts.Equals(t, test.expectErr, hasErr)
		})
	}
}

func TestGetAction(t *testing.T) {
	app, reg := appWithEmptyRegistration(mockConfig)
	def := definition.NewActionDefinition(&reg)

	instance := action.NewActionInstance(def.TypeName, mockConfig)
	action := reg.Constructor(instance, *reg.DefaultProps)
	expectedInstantiatedTypedAction := &InstantiatedTypedAction[ta.ActionValidEmpty]{
		Instance: instance,
		Action:   action,
	}

	actual, err := GetAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)

	asserts.Equals(t, expectedInstantiatedTypedAction, actual)
	asserts.Equals(t, nil, err)
}
