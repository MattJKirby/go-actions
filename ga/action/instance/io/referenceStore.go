package io


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