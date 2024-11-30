package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
)

type ModelInstance struct {
	ActionName string              `json:"name"`
	ActionUid  string              `json:"uid"`
	Parameters *parameter.Store    `json:"parameters"`
	Inputs     *io.Store[io.Input] `json:"inputs"`
}

type ModelInstanceConfig interface {
	GenerateUid() string
}

func NewModelInstance(typename string, config ModelInstanceConfig) *ModelInstance {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ModelInstance{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: parameter.NewStore(),
		Inputs:     io.NewStore[io.Input](ActionUid),
	}
}
