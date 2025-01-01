package action

type GoActionRegistration[Type GoAction] struct {
	Name string
	Constructor GoActionConstructor[Type]
}

type GoActionConstructor[Type GoAction] func(*ActionInstance) *Type

type GoAction interface {
	Execute()
}
