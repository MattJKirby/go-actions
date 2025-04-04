package store

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestWithUnsafeDecode(t *testing.T) {
	store := NewBaseStore(
		WithUnsafeUpdate[prop](true),
	)

	assert.Equals(t, true, store.config.unsafeUpdate)
}
