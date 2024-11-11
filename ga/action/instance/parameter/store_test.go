package parameter

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore()
	t.Run("test new store", func(t *testing.T) {
		asserts.Equals(t, 0, len(store.parameters))
	})
}

func TestGetOrDefault(t *testing.T) {
	store := NewStore()
	expectedGetParam := NewActionParameter("getParam", 0)
	expectedDefaultParam := NewActionParameter("defaultParam", 0)

	t.Run("test Get path", func(t *testing.T) {
		param := GetOrDefault("getParam", 0)(store)
		asserts.Equals(t, expectedGetParam, param)
	})

	t.Run("test Default path", func(t *testing.T) {
		param := GetOrDefault("defaultParam", 0)(store)
		asserts.Equals(t, expectedDefaultParam, param)
		asserts.Equals(t, &expectedDefaultParam, &param)
	})
}
