package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}
var mockActionConfig = &action.ActionConfig{}

func TestNewAction(t *testing.T) {
	reg := action.ActionRegistration[testActions.ActionValidEmpty]{Action: testActions.ActionValidEmpty{}}
	definition := definition.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockGlobalConfig, mockActionConfig)

	expected := &Action[action.GoAction]{
		Definition:       testActions.ActionValidEmpty{},
		Instance:         instance,
		BaseActionFields: NewBaseActionFields(instance),
	}

	actual, err := NewAction[action.GoAction](mockGlobalConfig, mockActionConfig, definition)
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, actual)
}
