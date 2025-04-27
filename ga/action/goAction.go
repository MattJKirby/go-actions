package action

type GoActionRegistration[Type GoAction] struct {
	Name         string
	Action       Type
}

type GoActionProps any

type GoAction interface {
	Init(*ActionInstance)
	Execute()
}

type TriggerAction interface {
	GoAction
	PublishTriggerConditions()
}
