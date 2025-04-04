package config

import (
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockUidGenerator = testHelpers.MockUidGenerator{MockUid: "uid"}

func TestWithCustomUidGenerator(t *testing.T) {
	config := DefaultGlobalConfig()
	WithCustomUidGenerator(mockUidGenerator)(config)
	assert.Equals(t, "uid", config.UidGenerator.GenerateUid())
}
