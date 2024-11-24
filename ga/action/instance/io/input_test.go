package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewInput(t *testing.T) {
	input := newInput("name", "actionUid")

	t.Run("test new input", func(t *testing.T) {
		asserts.Equals(t, "name", input.Name())
		asserts.Equals(t, "actionUid__Input:name", input.Id())
	})
}

// func TestGetOrDefault(t *testing.T){
// 	store := resourceStore.NewStore[Input]()

// 	t.Run("test default", func(t *testing.T) {
// 		expected := newInput("name", "id")
// 		input := GetOrDefault("input")(store)

// 		asserts.Equals(t, expected, input)

// 	})
// }
