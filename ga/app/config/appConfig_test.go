package config

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestWithGlobalConfig(t *testing.T) {
	cfg := DefaultApplicationConfig()
	WithGlobalConfigOptions(
		WithCustomUidGenerator(mockUidGenerator),
	)(cfg)

	assert.Equals(t, "uid", cfg.Global.UidGenerator.GenerateUid())
}
