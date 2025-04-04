package config

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestWithGlobalConfig(t *testing.T) {
	appConfig := NewAppConfig(
		WithGlobalConfig(
			WithCustomUidGenerator(mockUidGenerator),
		),
	)
	assert.Equals(t, "uid", appConfig.Global.UidGenerator.GenerateUid())
}
