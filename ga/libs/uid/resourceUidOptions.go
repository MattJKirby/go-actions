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

func WithUid(uid string) ResourceUidOption {
	return func(ru *ResourceUid) {
		ru.uid = uid
	}
}

func WithSubResource(resourceType string) ResourceUidOption {
	return func(ru *ResourceUid) {
		ru.subResourceType = resourceType
	}
}

func WithSubResourceId(resourceId string) ResourceUidOption {
	return func(ru *ResourceUid) {
		ru.subResourceId = resourceId
	}
}
