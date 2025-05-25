package executable

import (
	"go-actions/ga/action"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidProvider: mockGenerator}
var mockActionConfig = &action.ActionConfig{}

func TestNewAction(t *testing.T) {
	reg := action.ActionRegistration[testActions.ActionValidEmpty]{Action: testActions.ActionValidEmpty{}}
	definition := action.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance(mockGlobalConfig, mockActionConfig, definition)

	expected := &Action[action.GoAction]{
		Definition:       testActions.ActionValidEmpty{},
		Instance:         instance,
		BaseActionFields: NewBaseActionFields(instance),
	}

	actual, err := NewAction[action.GoAction](definition, instance)
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, actual)
}
