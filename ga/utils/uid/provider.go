package uid

import "github.com/google/uuid"

type Provider interface {
	New() string
}

type DefaultProvider struct{}

func (ug *DefaultProvider) New() string {
	return uuid.New().String()
}
