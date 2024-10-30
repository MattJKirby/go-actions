package parameter

type Parameter[T any] interface {
	Value() T
	DefaultValue() T
	SetValue(value T)
}

type ActionParameter[T any] struct {
	name string
	value T
	defaultValue T
}

func NewActionParameter[T any](Name string, DefaultValue T) *ActionParameter[T] {
	return &ActionParameter[T]{
		name: Name,
		value: DefaultValue,
		defaultValue: DefaultValue,
	}
}

func (ap *ActionParameter[T]) Value() T {
	return ap.value
}

func (ap *ActionParameter[T]) DefaultValue() T {
	return ap.defaultValue
}

func (ap *ActionParameter[T]) SetValue(value T) {
	ap.value = value
}
