package config

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestGenerateUid(t *testing.T){
	config := NewInstanceConfig()
	t.Run("default config", func(t *testing.T) {
		if config.GenerateUid() == "" {
			t.Errorf("uuidGenerator can't be empty string")
		}
		asserts.Equals(t, 36, len(config.GenerateUid()))
	})
}