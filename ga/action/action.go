package action

// type ActionConstructor = func () ActionConstructorValue

// type ActionConstructorValue interface {
// 	Execute()
// }

type Action struct {
	def *ActionDefinition
}

func NewAction(def *ActionDefinition) *Action {
	
	return &Action{
		def,
	}
}