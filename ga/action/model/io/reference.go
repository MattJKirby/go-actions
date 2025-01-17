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

func AssignReferences(in *Input, out *Output) {
	inputRef := NewReference(in.actionUid, in.Name)
	outputRef := NewReference(out.actionUid, out.Name)

	in.AssignOutput(outputRef)
	out.AssignInputReference(inputRef)
}
