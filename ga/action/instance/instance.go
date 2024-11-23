package instance

import (
	"fmt"
	"go-actions/ga/action/instance/io/references"
	"go-actions/ga/action/instance/parameter"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string           `json:"name"`
	ActionUid  string           `json:"uid"`
	Parameters *parameter.Store `json:"parameters"`
	Inputs *references.Store[references.Input] `json:"inputs"`
	Outputs *references.Store[references.Output] `json:"outputs"`
}

func NewActionInstance(typename string) *ActionInstance {
	uid := fmt.Sprintf("Action:%s:%s", typename, uuid.New())
	return &ActionInstance{
		ActionName: typename,
		ActionUid:  uid,
		Parameters: parameter.NewStore(),
		Inputs: references.NewActionReferenceStore[references.Input](),
		Outputs: references.NewActionReferenceStore[references.Output](),
	}
}
