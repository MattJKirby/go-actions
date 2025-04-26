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

func TestNewAction(t *testing.T) {
	reg := testActions.GenerateActionValidEmptyRegistration()
	definition := definition.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockGlobalConfig)

	expected := &Action[action.GoAction]{
		Action:   testActions.ActionValidEmpty{},
		Instance: instance,
	}

	actual, err := NewAction[action.GoAction](mockGlobalConfig, definition)
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, actual)
}
