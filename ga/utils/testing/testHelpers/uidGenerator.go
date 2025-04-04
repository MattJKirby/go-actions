package testHelpers

type MockUidGenerator struct {
	MockUid string
}

func (mac *MockUidGenerator) GenerateUid() string {
	return mac.MockUid
}
