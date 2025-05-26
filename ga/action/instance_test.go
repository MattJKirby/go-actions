package action_test

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/testHelpers"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidProvider: mockGenerator}
