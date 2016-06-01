package gull

import (
	"testing"

	"github.com/c2fo/gull/source/lib/gull/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (suite *ConfigTestSuite) TestParseJSONString() {
	result, err := NewConfigFromJson(testdata.ValidJsonConfig1, false)

	assert.Nil(suite.T(), err)
	leaf, err := result.Leaves.GetValue("/default/enableLogging")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "false", leaf)
}

func (suite *ConfigTestSuite) TestParseJSONArray() {
	result, err := NewConfigFromJson(testdata.ValidJsonConfig1, false)

	assert.Nil(suite.T(), err)
	leaf, err := result.Leaves.GetValue("/default/services")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "[well hi there]", leaf)
}

func (suite *ConfigTestSuite) TestInvalidPath() {
	result, err := NewConfigFromJson(testdata.ValidJsonConfig1, false)

	assert.Nil(suite.T(), err)
	_, err = result.Leaves.GetValue("/invalid/services")

	assert.NotNil(suite.T(), err)
}
