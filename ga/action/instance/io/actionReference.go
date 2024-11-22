package io

import "fmt"

type ActionReference struct {
	actionId string
	referenceName string
	id string
}

func NewActionReference(actionId string, referenceId string, referenceType string) *ActionReference {
	id := fmt.Sprintf("%s__ref:%s:%s", actionId, referenceType, referenceId)
	return &ActionReference{
		actionId,
		referenceId,
		id,
	}
}

type ActionOutputReference struct {
	ActionReference
	outputId string
}

func NewActionOutputReference(actionId string, outputId string) *ActionOutputReference {
	return &ActionOutputReference{
		ActionReference: *NewActionReference(actionId, outputId, "output"),
		outputId: outputId,
	}
}

type ActionInputReference struct {
	ActionReference
	inputId string
}

func NewActionInputReference(actionId string, inputId string) *ActionInputReference {
	return &ActionInputReference{
		ActionReference: *NewActionReference(actionId, inputId, "input"),
		inputId: inputId,
	}
}
