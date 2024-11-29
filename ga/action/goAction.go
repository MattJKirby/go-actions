package action

type GoActionConstructor[T GoAction] func(GoActionInternals) *T

type GoAction interface {
	Execute()
}
