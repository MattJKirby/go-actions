package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
)

type ActionModel struct {
	ActionName string                `json:"name"`
	ActionUid  string                `json:"uid"`
	Parameters *parameter.Store      `json:"parameters"`
	Inputs     *io.IOStore[io.Input] `json:"inputs"`
}

type ActionModelConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionModelConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: parameter.NewStore(),
		Inputs:     io.NewIOStore[io.Input](ActionUid),
	}
}
