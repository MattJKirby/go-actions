package io

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/testHelpers"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

