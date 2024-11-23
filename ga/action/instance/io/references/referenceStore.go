package references


type ActionReferenceStore[T ActionReference] struct {
	references map[string]*T
}

func NewActionReferenceStore[T ActionReference]() *ActionReferenceStore[T] {
	return &ActionReferenceStore[T]{
		references: make(map[string]*T),
	}
}

func (ars *ActionReferenceStore[T]) Add(ref *T) {
	ars.references[(*ref).Id()] = ref
}

func (ars *ActionReferenceStore[T]) GetOrDefault(ref *T) *T{
	refId := (*ref).Id()
	_, exists := ars.references[refId]
	if !exists {
		ars.Add(ref)
	}

	return ref
}