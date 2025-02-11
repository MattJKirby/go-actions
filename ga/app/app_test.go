package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
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
	def := definition.NewActionDefinition(&reg)
	actual := executable.NewExecutableAction(mockConfig, def.ActionTypeDefinition)

	tests := []struct {
		name               string
		inputName          string
		expectedExecutable *executable.ExecutableAction
		expectErr          bool
	}{
		{name: "valid - existing action name", inputName: "ActionValidEmpty", expectedExecutable: actual, expectErr: false},
		{name: "invalid - not existing action name", inputName: "notregistered", expectedExecutable: nil, expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := InstantiateAction(test.inputName)(app)
			hasErr := err != nil

			asserts.Equals(t, test.expectedExecutable, actual)
			asserts.Equals(t, test.expectErr, hasErr)
		})
	}
}

func TestInstantiateTypedAction(t *testing.T) {
	app, reg := appWithEmptyRegistration(mockConfig)
	def := definition.NewActionDefinition(&reg)

	expectedExecutable := executable.NewExecutableAction(mockConfig, def.ActionTypeDefinition)
	expectedAction, _ := expectedExecutable.Action.(*ta.ActionValidEmpty)
	expectedTypedExecutable := &executable.TypedExecutable[ta.ActionValidEmpty, ta.ActionValidEmptyProps]{
		ExecutableAction: expectedExecutable,
		Action:           expectedAction,
	}

	actual, err := InstantiateTypedAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)

	asserts.Equals(t, expectedTypedExecutable, actual)
	asserts.Equals(t, nil, err)
}
