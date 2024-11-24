package input

import (
	"go-actions/ga/action/instance/io"
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

func TestGetOrDefault(t *testing.T) {
	store := io.NewStore[Input]("uid")

	t.Run("test default", func(t *testing.T) {
		expected := newInput("name", "uid")
		input := GetOrDefaultInput("name")(store)

		asserts.Equals(t, expected, input)

	})
}
