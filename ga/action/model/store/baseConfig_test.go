package store

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestWithUnsafeDecode(t *testing.T) {
	store := NewBaseStore(
		WithUnsafeDecode[prop](true),
	)

	asserts.Equals(t, true, store.config.unsafeUpdate)
}
