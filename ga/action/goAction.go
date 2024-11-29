package action

import (
	"go-actions/ga/action/internals"
)

type GoActionConstructor[T GoAction] func(internals.GoActionInternals) *T

type GoAction interface {
	Execute()
}
