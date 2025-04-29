package app

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockUidGenerator = testHelpers.MockUidGenerator{MockUid: "uid"}

func TestWithGlobalConfig(t *testing.T) {
	cfg := DefaultApplicationConfig()
	WithGlobalConfigOptions(
		config.WithCustomUidGenerator(mockUidGenerator),
	)(cfg)

	assert.Equals(t, "uid", cfg.Global.UidGenerator.GenerateUid())
}
