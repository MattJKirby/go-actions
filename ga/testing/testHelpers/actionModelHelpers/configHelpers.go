package actionmodelhelpers

type MockActionModelConfig struct {
	mockUid string
}

func (mamc *MockActionModelConfig) GenerateUid() string {
	return mamc.mockUid
}