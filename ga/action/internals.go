package action

type ActionInternals struct {
	Instance *ActionInstance
}

func NewActionInternals(instance *ActionInstance) *ActionInternals {
	return &ActionInternals{
		Instance: instance,
	}
}