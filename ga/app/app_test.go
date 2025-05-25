package app

import (
	"go-actions/ga/action"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"

	"testing"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidProvider: mockGenerator}
var mockAppConfig = &ApplicationConfig{Global: mockGlobalConfig}

func appWithEmptyRegistration() (*App, action.ActionRegistration[ta.ActionValidEmpty]) {
	app := NewApp("test")
	app.Config = mockAppConfig
	reg := action.ActionRegistration[ta.ActionValidEmpty]{Action: ta.ActionValidEmpty{}}
	RegisterAction(&reg)(app)
	return app, reg
}

func TestRegisterActionAndGet(t *testing.T) {
	app, _ := appWithEmptyRegistration()
	result, _ := GetDefinitionByType[ta.ActionValidEmpty]()(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestGetDefinitionByName(t *testing.T) {
	app, reg := appWithEmptyRegistration()
	result, err := GetDefinitionByName("ActionValidEmpty")(app)

	expectedTypeDef := action.TypeDefinitionFromRegistration(&reg)

	assert.Equals(t, expectedTypeDef, result)
	assert.Equals(t, nil, err)
}

func TestAppGetActionByName(t *testing.T) {
	app, _ := appWithEmptyRegistration()

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
			_, err := GetActionByName(test.inputName, nil)(app)
			hasErr := err != nil
			assert.Equals(t, test.expectErr, hasErr)
		})
	}
}

func TestGetAction(t *testing.T) {
	app, _ := appWithEmptyRegistration()

	_, err := GetAction[ta.ActionValidEmpty]()(app)

	assert.Equals(t, nil, err)
}
