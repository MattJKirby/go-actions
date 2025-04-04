package uid

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestGenerateUid(t *testing.T) {
	generator := &DefaultUidGenerator{}

	assert.Equals(t, 36, len(generator.GenerateUid()))
}
