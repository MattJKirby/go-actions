package actionModelTestHelpers

type MockActionModelConfig struct {
	MockUid string
}

func (mamc *MockActionModelConfig) GenerateUid() string {
	return mamc.MockUid
}