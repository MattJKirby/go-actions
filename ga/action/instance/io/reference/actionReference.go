package reference

import "fmt"


type actionReference struct {
	actionId   string
	resourceId string
	id         string
}

func newReference(actionId string, resourceId string, referenceType string) *actionReference {
	id := fmt.Sprintf("%s__ref:%s:%s", actionId, referenceType, resourceId)
	return &actionReference{
		actionId,
		resourceId,
		id,
	}
}


type Output struct {
	actionReference
	outputId string
}

func NewOutput(actionId string, outputId string) *Output {
	return &Output{
		actionReference: *newReference(actionId, outputId, "output"),
		outputId:        outputId,
	}
}

type Input struct {
	actionReference
	inputId string
}

func NewInput(actionId string, inputId string) *Input {
	return &Input{
		actionReference: *newReference(actionId, inputId, "input"),
		inputId:         inputId,
	}
}
