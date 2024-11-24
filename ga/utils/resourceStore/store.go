package resourceStore

type Store[T any] struct {
	resources map[string]*T
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{
		resources: make(map[string]*T),
	}
}

func (s *Store[T]) GetOrDefault(name string, resource *T) *T {
	_, exists := s.resources[name]
	if !exists {
		s.resources[name] = resource
	}

	return s.resources[name]
}
