package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewActionReference(t *testing.T) {
	ref := NewActionReference("actionId", "referenceId", "refType")
	
	t.Run("assert reference id", func(t *testing.T) {
		asserts.Equals(t, "actionId__ref:refType:referenceId", ref.id)
	})
}

func TestNewActionOutputReference(t *testing.T) {
	outputRef := NewActionOutputReference("actionId", "outputId")
	t.Run("assert output id", func(t *testing.T) {
		asserts.Equals(t, "outputId", outputRef.outputId)
	})
}

func TestNewActionInputReference(t *testing.T) {
	inputRef := NewActionInputReference("actionId", "inputId")
	t.Run("assert input id", func(t *testing.T) {
		asserts.Equals(t, "inputId", inputRef.inputId)
	})
}