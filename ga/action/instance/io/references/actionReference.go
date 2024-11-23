package references

import "fmt"

type ActionReference interface {
	Id() string
}

type actionReference struct {
	actionId string
	recieverId string
	id string
}

func newActionReference(actionId string, recieverId string, referenceType string) *actionReference {
	id := fmt.Sprintf("%s__ref:%s:%s", actionId, referenceType, recieverId)
	return &actionReference{
		actionId,
		recieverId,
		id,
	}
}

func (ar actionReference) Id() string {
	return ar.id
}

type ActionOutputReference struct {
	actionReference
	outputId string
}

func NewActionOutputReference(actionId string, outputId string) *ActionOutputReference {
	return &ActionOutputReference{
		actionReference: *newActionReference(actionId, outputId, "output"),
		outputId: outputId,
	}
}

type ActionInputReference struct {
	actionReference
	inputId string
}

func NewActionInputReference(actionId string, inputId string) *ActionInputReference {
	return &ActionInputReference{
		actionReference: *newActionReference(actionId, inputId, "input"),
		inputId: inputId,
	}
}
