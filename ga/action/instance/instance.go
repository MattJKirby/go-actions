package instance

import (
	"fmt"
	"go-actions/ga/action/instance/parameter"

	"github.com/google/uuid"
)

type ActionInstance struct {
	ActionName string
	ActionUid  string
	Parameters *parameter.Store
}

func NewActionInstance(typename string) *ActionInstance {
	uid := fmt.Sprintf("Action:%s:%s", uuid.New(), typename)
	return &ActionInstance{
		ActionName: typename,
		ActionUid:  uid,
		Parameters: parameter.NewStore(),
	}
}
