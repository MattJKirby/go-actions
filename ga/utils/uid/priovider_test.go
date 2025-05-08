package uid

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestGenerateUid(t *testing.T) {
	generator := &DefaultProvider{}

	assert.Equals(t, 36, len(generator.New()))
}
