package store

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestWithUnsafeDecode(t *testing.T) {
	store := NewBaseStore(
		WithUnsafeDecode[prop](),
	)

	asserts.Equals(t, true, store.config.unsafeDecode)
}
