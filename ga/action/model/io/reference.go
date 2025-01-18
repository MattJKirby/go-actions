package io

type ActionReference struct {
	ActionUid    string `json:"actionUid"`
	ResourceName string `json:"resourceName"`
}

func NewReference(ActionUid string, ResourceName string) *ActionReference {
	return &ActionReference{
		ActionUid,
		ResourceName,
	}
}

func AssignReferences(source *Output, targets []*Input) {
	sourceRef := NewReference(source.actionUid, source.Name)

	for _, target := range targets {
		targetRef := NewReference(target.actionUid, target.Name)
		source.AssignTarget(targetRef)
		target.AssignSource(sourceRef)
	}
}
