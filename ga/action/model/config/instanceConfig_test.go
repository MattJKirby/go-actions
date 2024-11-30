package config

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestGenerateUid(t *testing.T) {
	config := NewInstanceConfig()
	t.Run("default config uid", func(t *testing.T) {
		asserts.Equals(t, 36, len(config.GenerateUid()))
	})
}
