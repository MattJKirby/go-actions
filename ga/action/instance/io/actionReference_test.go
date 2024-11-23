package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewActionReference(t *testing.T) {
	ref := newActionReference("actionId", "recieverId", "refType")
	
	t.Run("assert reference id", func(t *testing.T) {
		asserts.Equals(t, "actionId__ref:refType:recieverId", ref.id)
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