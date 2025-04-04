package actionTestHelpers

type MockActionConfig struct {
	MockUid string
}

func (mac *MockActionConfig) GenerateUid() string {
	return mac.MockUid
}
