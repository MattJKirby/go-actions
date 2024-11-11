package parameter

type Store struct {
	parameters map[string]any
}

func NewStore() *Store {
	return &Store{
		parameters: make(map[string]any),
	}
}

func GetOrDefault[T any](name string, defaultValue T) func(*Store) *ActionParameter[T] {
	return func(s *Store) *ActionParameter[T] {
		_, exists := s.parameters[name]
		if !exists {
			s.parameters[name] = any(NewActionParameter(name, defaultValue))
		}

		return any(s.parameters[name]).(*ActionParameter[T])
	}
}
