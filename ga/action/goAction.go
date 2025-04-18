package action

type GoActionRegistration[Type GoAction, Props GoActionProps] struct {
	Name         string
	Action       Type
	DefaultProps Props
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
