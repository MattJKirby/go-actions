package instance

import (
	"fmt"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string
	ActionUid string
}

func NewActionInstance(typename string) *ActionInstance {
	uid := fmt.Sprintf("Action:%s:%s", uuid.New(), typename)
	return &ActionInstance{
		ActionName: typename,
		ActionUid: uid,
	}
}