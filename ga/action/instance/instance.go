package instance

import (
	"fmt"
	reference "go-actions/ga/action/instance/io/references"
	"go-actions/ga/action/instance/parameter"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string                             `json:"name"`
	ActionUid  string                             `json:"uid"`
	Parameters *parameter.Store                   `json:"parameters"`
	Inputs     *reference.Store[reference.Input]  `json:"inputs"`
	Outputs    *reference.Store[reference.Output] `json:"outputs"`
}

func NewActionInstance(typename string) *ActionInstance {
	uid := fmt.Sprintf("Action:%s:%s", typename, uuid.New())
	return &ActionInstance{
		ActionName: typename,
		ActionUid:  uid,
		Parameters: parameter.NewStore(),
		Inputs:     reference.NewActionReferenceStore[reference.Input](),
		Outputs:    reference.NewActionReferenceStore[reference.Output](),
	}
}
