package parameter

import (
	"go-actions/ga/cr/asserts"
	"reflect"
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

func TestGet(t *testing.T) {
	store := NewStore()
	expectedGetParam := NewActionParameter("intParam", 0)
	GetOrDefault("intParam", 0)(store)

	t.Run("test get with meta", func(t *testing.T) {
		param, _ := store.Get("intParam")
		asserts.Equals(t, reflect.TypeOf(0), param.parameterType)
		asserts.Equals(t, any(expectedGetParam), param.parameterValue)
	})

	t.Run("test no such parameter", func(t *testing.T) {
		param, err := store.Get("bad")
		if err == nil {
			t.Errorf("expected error but got nil")
		}

		if param != nil {
			t.Errorf("expected nil but got %v", param)
		}
	})

}