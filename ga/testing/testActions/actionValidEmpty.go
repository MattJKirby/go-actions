package testActions

import (
	"fmt"
	"go-actions/ga/action"
)

type ActionValidEmpty struct{}

type ActionValidEmptyProps struct{}

func (tav ActionValidEmpty) Execute() {
	fmt.Println("Executing ActionValidEmpty")
}

func GenerateActionValidEmptyCtor() action.GoActionConstructor[ActionValidEmpty, ActionValidEmptyProps] {
	return GenerateEmptyCtor[ActionValidEmpty, ActionValidEmptyProps]()
}

func GenerateActionValidEmptyRegistration() action.GoActionRegistration[ActionValidEmpty, ActionValidEmptyProps] {
	ctor := GenerateActionValidEmptyCtor()
	return GenerateRegistration(ctor, &ActionValidEmptyProps{})
}
