package uid

type ResourceUidOption func(*ResourceUid)

func WithNamespace(namespace string) ResourceUidOption {
	return func(ru *ResourceUid) {
		ru.namespace = namespace
	}
}

func WithResource(resource string) ResourceUidOption {
	return func(ru *ResourceUid) {
		ru.resource = resource
	}
}
