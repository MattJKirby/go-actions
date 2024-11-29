package instance

import (
	"fmt"
	"go-actions/ga/action/instance/io"
	"go-actions/ga/action/instance/parameter"
)

type ActionInstance struct {
	ActionName string              `json:"name"`
	ActionUid  string              `json:"uid"`
	Parameters *parameter.Store    `json:"parameters"`
	Inputs     *io.Store[io.Input] `json:"inputs"`
}

type ActionInstanceConfig interface {
	GenerateUid() string
}

func NewActionInstance(typename string, config ActionInstanceConfig) *ActionInstance {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionInstance{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: parameter.NewStore(),
		Inputs:     io.NewStore[io.Input](ActionUid),
	}
}
