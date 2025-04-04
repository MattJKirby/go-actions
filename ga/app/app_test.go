package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers/actionTestHelpers"

	"testing"
)

var mockConfig = &actionTestHelpers.MockActionConfig{MockUid: "uid"}

func appWithEmptyRegistration(config model.ActionConfig) (*App, action.GoActionRegistration[ta.ActionValidEmpty, ta.ActionValidEmptyProps]) {
	app := NewApp("test")
	app.actionConfig = config
	registration := ta.GenerateActionValidEmptyRegistration()
	RegisterAction(&registration)(app)
	return app, registration
}

func TestRegisterActionAndGet(t *testing.T) {
	app, _ := appWithEmptyRegistration(mockConfig)
	result, _ := GetDefinitionByType[ta.ActionValidEmpty, ta.ActionValidEmptyProps]()(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestGetDefinitionByName(t *testing.T) {
	app, reg := appWithEmptyRegistration(mockConfig)
	result, err := GetDefinitionByName("ActionValidEmpty")(app)

	expectedTypeDef := definition.TypeDefinitionFromRegistration(&reg)

	assert.Equals(t, expectedTypeDef, result)
	assert.Equals(t, nil, err)
}

func TestAppGetActionByName(t *testing.T) {
	app, _ := appWithEmptyRegistration(mockConfig)

	tests := []struct {
		name      string
		inputName string
		expectErr bool
	}{
		{name: "valid - existing action name", inputName: "ActionValidEmpty", expectErr: false},
		{name: "invalid - not existing action name", inputName: "notregistered", expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := GetActionByName(test.inputName)(app)
			hasErr := err != nil
			assert.Equals(t, test.expectErr, hasErr)
		})
	}
}

func TestGetAction(t *testing.T) {
	app, _ := appWithEmptyRegistration(mockConfig)

	_, err := GetAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)

	assert.Equals(t, nil, err)
}
