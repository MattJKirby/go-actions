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