package reference

type actionReference struct {
	ActionUid    string `json:"actionUid"`
	ResourceName string `json:"resourceName"`
	referenceType string
}

func newReference(ActionUid string, ResourceName string, referenceType string) *actionReference {
	return &actionReference{
		ActionUid,
		ResourceName,
		referenceType,
	}
}

type OutputReference struct {
	*actionReference
}

func NewOutputReference(ActionUid string, outputName string) *OutputReference {
	return &OutputReference{
		actionReference: newReference(ActionUid, outputName, "output"),
	}
}
