package instance

import (
	"fmt"
	"go-actions/ga/action/instance/io"
	"go-actions/ga/action/instance/parameter"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string              `json:"name"`
	ActionUid  string              `json:"uid"`
	Parameters *parameter.Store    `json:"parameters"`
	Inputs     *io.Store[io.Input] `json:"inputs"`
}

func NewActionInstance(typename string) *ActionInstance {
	uid := fmt.Sprintf("%s:%s", typename, uuid.New())
	return &ActionInstance{
		ActionName: typename,
		ActionUid:  uid,
		Parameters: parameter.NewStore(),
		Inputs:     io.NewStore[io.Input](uid),
	}
}
