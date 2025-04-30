package action

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockUidGenerator = &testHelpers.MockUidGenerator{MockUid: "abc"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockUidGenerator}

func TestUid(t *testing.T) {
	mockActionConfig := &ActionConfig{UidFormat: `uid:%s:%s`}
	uid := NewActionUid(mockGlobalConfig, mockActionConfig, "actionName")

	assert.Equals(t, "actionName", uid.actionName)
	assert.Equals(t, mockUidGenerator.MockUid, uid.uid)
	assert.Equals(t, "uid:actionName:abc", uid.GetUid())
}