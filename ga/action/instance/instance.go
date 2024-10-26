package instance

import (
	"fmt"
	"go-actions/ga/action/definition"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string
	ActionUid string
}

func NewActionInstance(def *definition.ActionDefinition) *ActionInstance {
	uid := fmt.Sprintf("Action:%s:%s", uuid.New(), def.Name)
	return &ActionInstance{
		ActionName: def.Name,
		ActionUid: uid,
	}
}