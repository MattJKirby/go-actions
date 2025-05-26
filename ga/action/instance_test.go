package action_test

import (
	"go-actions/ga/action"
	"go-actions/ga/action/model"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"

	"testing"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidProvider: mockGenerator}

func TestNewActionInstance(t *testing.T) {
	def := &action.TypeDefinition{TypeName: "test"}

	instance := action.NewActionInstance(mockConfig, def)
	expectedModel := model.NewActionModel(mockConfig, instance.Uid)

	assert.Equals(t, expectedModel, instance.Model)
}
