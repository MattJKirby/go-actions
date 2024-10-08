package action

import (
	"fmt"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string
	ActionUid string
}

func NewActionInstance(def *ActionDefinition) *ActionInstance {
	uid := fmt.Sprintf("Action:%s:%s", uuid.New(), def.typeName)
	return &ActionInstance{
		ActionName: def.typeName,
		ActionUid: uid,
	}
}