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

// func AssignReferences(in *Input, out *Output){
// 	in.AssignOutput(out.)
// }