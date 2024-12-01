package config

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestGenerateUid(t *testing.T) {
	config := NewModelConfig()

	asserts.Equals(t, 36, len(config.GenerateUid()))
}
