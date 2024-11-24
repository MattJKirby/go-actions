package input

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action/instance/io"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewInput(t *testing.T) {
	input := newInput("name", "actionUid")

	z,_ := json.Marshal(input)
	fmt.Println(string(z))

	t.Run("test new input", func(t *testing.T) {
		asserts.Equals(t, "name", input.name)
		asserts.Equals(t, "actionUid__Input:name", input.id)
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

// func TestCustomMarshalling(t *testing.T){
// 	input := newInput("name", "actionUid")

// 	t.Run("custom marshalling", func(t *testing.T) {
// 		asserts.Equals()
// 	})
// }
