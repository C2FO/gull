package ui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LaunchTestSuite struct {
	suite.Suite
}

func TestLaunchSuite(t *testing.T) {
	testSuite := &LaunchTestSuite{}
	suite.Run(t, testSuite)
}

func (suite *LaunchTestSuite) TestInvalidArguments() {
	defer func() {
		err := recover()
		assert.Nil(suite.T(), err)
	}()
	os.Args = []string{"up", "--doesntexist"}
	Launch()
}
