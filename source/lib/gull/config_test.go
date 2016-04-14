package gull

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IngestTestSuite struct {
	suite.Suite
	JsonSample1 string
}

func TestIngestSuite(t *testing.T) {
	testSuite := IngestTestSuite{
		JsonSample1: "{\"*\":{\"services\":[\"well\",\"hi\",\"there\"],\"enableLogging\":false},\"production\": {\"enableLogging\":true}}",
	}
	suite.Run(t, &testSuite)
}

func (suite *IngestTestSuite) TestParseJSONString() {
	result, err := NewConfigFromJson(suite.JsonSample1)

	assert.Nil(suite.T(), err)
	leaf, err := result.GetPath("/default/enableLogging")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "false", leaf)
}

func (suite *IngestTestSuite) TestParseJSONArray() {
	result, err := NewConfigFromJson(suite.JsonSample1)

	assert.Nil(suite.T(), err)
	leaf, err := result.GetPath("/default/services")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "[well hi there]", leaf)
}

func (suite *IngestTestSuite) TestInvalidPath() {
	result, err := NewConfigFromJson(suite.JsonSample1)

	assert.Nil(suite.T(), err)
	_, err = result.GetPath("/invalid/services")

	assert.NotNil(suite.T(), err)
}
