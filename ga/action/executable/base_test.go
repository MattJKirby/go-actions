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

func TestNewBaseExecutable(t *testing.T) {
	reg := testActions.GenerateActionValidEmptyRegistration()
	definition := definition.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockGlobalConfig)

	expected := &baseExecutable{
		construct: reg.Constructor(instance, testActions.ActionValidEmptyProps{}),
		instance:  instance,
	}

	actual, err := newBaseExecutable(mockGlobalConfig, definition)
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, actual)
}
