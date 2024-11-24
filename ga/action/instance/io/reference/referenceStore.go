package reference

type Store[T ActionReference] struct {
	references map[string]*T
}

func NewActionReferenceStore[T ActionReference]() *Store[T] {
	return &Store[T]{
		references: make(map[string]*T),
	}
}

func (ars *Store[T]) Add(ref *T) {
	ars.references[(*ref).Id()] = ref
}

func (ars *Store[T]) GetOrDefault(ref *T) *T {
	refId := (*ref).Id()
	_, exists := ars.references[refId]
	if !exists {
		ars.Add(ref)
	}

	return ref
}
