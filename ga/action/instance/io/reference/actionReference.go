package reference

type ActionReference struct {
	ActionUid    string `json:"actionUid"`
	ResourceName string `json:"resourceName"`
}

func NewReference(ActionUid string, ResourceName string, referenceType string) *ActionReference {
	return &ActionReference{
		ActionUid,
		ResourceName,
	}
}
