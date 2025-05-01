package action

type ActionUidBuilder interface {
	GetActionUid() string
	GetResourceUid(string, string) string
}

type ActionConfig struct {
	UidFormat  string
	UidBuilder *ActionUidBuilder
}

func DefaultActionConfig() *ActionConfig {
	return &ActionConfig{
		UidFormat:  `ga:%s:%s:%s:%s`,
		UidBuilder: nil,
	}
}
