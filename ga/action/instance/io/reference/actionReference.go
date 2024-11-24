package reference

import "fmt"

type ActionReference struct {
	ActionUid    string `json:"actionUid"`
	ResourceName string `json:"resourceName"`
	id           string
}

func NewReference(ActionUid string, ResourceName string, referenceType string) *ActionReference {
	id := fmt.Sprintf("%s__ref:%s:%s", ActionUid, referenceType, ResourceName)
	return &ActionReference{
		ActionUid,
		ResourceName,
		id,
	}
}
