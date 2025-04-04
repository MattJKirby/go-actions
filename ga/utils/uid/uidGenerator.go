package uid

import "github.com/google/uuid"

type UidGenerator interface {
	GenerateUid() string
}

type DefaultUidGenerator struct {}

func (ug *DefaultUidGenerator) GenerateUid() string {
	return uuid.New().String()
}