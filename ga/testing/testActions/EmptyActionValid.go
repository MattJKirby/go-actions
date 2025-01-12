package testActions

import "go-actions/ga/action"

type EmptyActionValid struct{}

type EmptyActionValidProps struct{}

func (tav EmptyActionValid) Execute() {}

func GenerateEmptyActionValidCtor() action.GoActionConstructor[EmptyActionValid, EmptyActionValidProps] {
	return GenerateEmptyCtor[EmptyActionValid, EmptyActionValidProps]()
}

func GenerateEmptyActionValidRegistration() action.GoActionRegistration[EmptyActionValid, EmptyActionValidProps] {
	ctor := GenerateEmptyActionValidCtor()
	return GenerateRegistration(ctor, &EmptyActionValidProps{})
}
