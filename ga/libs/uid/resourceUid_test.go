package uid

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockUidGenerator = &testHelpers.MockUidGenerator{MockUid: "abc"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockUidGenerator}

func TestGetString(t *testing.T) {
	uid := NewResourceUid(mockGlobalConfig, WithResource("someAction"))
	assert.Equals(t, "ga:core:someaction:abc::", uid.GetString())
}

func TestGetSecondary(t *testing.T) {
	uid := NewResourceUid(mockGlobalConfig, WithResource("someAction"))
	assert.Equals(t, "ga:core:someaction:abc:a:b", uid.GetSecondaryUid("a", "b"))
}