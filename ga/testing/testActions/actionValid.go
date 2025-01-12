package testActions

import "go-actions/ga/action"

type ActionValid struct{}

type ActionValidProps struct{
	Prop string
}

func (tav ActionValid) Execute() {}

func GenerateActionValidCtor() action.GoActionConstructor[ActionValid, ActionValidProps] {
	return GenerateEmptyCtor[ActionValid, ActionValidProps]()
}

func GenerateActionValidRegistration() action.GoActionRegistration[ActionValid, ActionValidProps] {
	ctor := GenerateActionValidCtor()
	return GenerateRegistration(ctor, &ActionValidProps{})
}
