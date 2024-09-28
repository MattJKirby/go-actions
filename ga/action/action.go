package action

type Constructor[T FunctionDefinition] func () T

type FunctionDefinition interface {
	Execute()
}

type Action struct {
	def *ActionDefinition
}

func NewAction(def *ActionDefinition) *Action {

	return &Action{
		def,
	}
}
