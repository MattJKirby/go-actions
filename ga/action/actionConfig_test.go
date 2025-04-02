package action

import (
	"go-actions/ga/testing/assert"
	"testing"
)

func TestGenerateUid(t *testing.T) {
	config := NewConfig()

	assert.Equals(t, 36, len(config.GenerateUid()))
}
