package action

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestGenerateUid(t *testing.T) {
	config := NewConfig()

	asserts.Equals(t, 36, len(config.GenerateUid()))
}
