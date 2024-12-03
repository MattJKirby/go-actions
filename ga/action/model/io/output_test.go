package io

import (
	"go-actions/ga/action/model/io/reference"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewOutput(t *testing.T) {
	output := NewActionOutput("name", "uid")
	asserts.Equals(t, "name", output.Name)
	asserts.Equals(t, "uid__Output:name", output.Id)
	asserts.Equals(t, 0, len(output.InputReferences))
}

func TestAssignInputRef(t *testing.T) {
	output := NewActionOutput("name", "uid")
	inputRef := reference.NewInputReference("other uid", "output")

	output.AssignInputReference(inputRef)
	asserts.Equals(t, 1, len(output.InputReferences))
	asserts.Equals(t, inputRef, output.InputReferences[0])
}