package action_test

import (
	"go-actions/ga/action"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"testing"
)

// var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
// var mockGlobalConfig = &config.GlobalConfig{UidProvider: mockGenerator}

func TestNewAction(t *testing.T) {
	reg := action.ActionRegistration[testActions.ActionValidEmpty]{Action: testActions.ActionValidEmpty{}}
	definition := action.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance(mockGlobalConfig, definition)

	actual, err := action.NewAction[testActions.ActionValidEmpty](definition, instance)
	assert.Equals(t, nil, err)
	assert.Equals(t, testActions.ActionValidEmpty{}, actual.Definition)
	assert.Equals(t, actual.Instance, instance)
}
