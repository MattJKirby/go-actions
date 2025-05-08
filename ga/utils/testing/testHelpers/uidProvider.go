package testHelpers

type MockUidProvider struct {
	MockUid string
}

func (mac MockUidProvider) New() string {
	return mac.MockUid
}
