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

func TestGetActionSuccessfulNilProps(t *testing.T) {
	app, _ := appWithEmptyRegistration(mockConfig)
	_, err := GetTypedAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)

	asserts.Equals(t, false, err != nil)
}

func TestGetActionWithProps(t *testing.T) {
	app := NewApp("test")
	reg := ta.GenerateActionValidRegistration()
	RegisterAction(&reg)(app)

	props := &ta.ActionValidProps{Param1: "asdf"}
	act, err := GetTypedAction[ta.ActionValid](props)(app)

	asserts.Equals(t, false, err != nil)
	asserts.Equals(t, props.Param1, act.Action.Param1.Value())

}

func TestGetActionFail(t *testing.T) {
	app := NewApp("test")

	_, err := GetTypedAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
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
