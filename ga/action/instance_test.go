package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"

	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}
var actionConfig = &ActionConfig{}

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("test", mockConfig, actionConfig)
	expectedModel := model.NewActionModel("test", mockConfig)

	assert.Equals(t, expectedModel, instance.Model)
}
