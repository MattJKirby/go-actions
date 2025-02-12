package action

type GoActionRegistration[Type GoAction, Props GoActionProps] struct {
	Name         string
	Constructor  GoActionConstructor[Type, Props]
	DefaultProps *Props
}

type GoActionProps any

type GoActionConstructor[Type GoAction, Props GoActionProps] func(*ActionInstance, Props) *Type

type GoAction interface {
	Execute()
}
