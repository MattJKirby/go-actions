package config

import (
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockUidGenerator = testHelpers.MockUidProvider{MockUid: "uid"}

func TestWithCustomUidGenerator(t *testing.T) {
	config := DefaultGlobalConfig()
	WithCustomUidProvider(mockUidGenerator)(config)
	assert.Equals(t, "uid", config.UidProvider.New())
}
