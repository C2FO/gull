package gull

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/c2fo/gull/source/lib/gull/testdata"
)

type ConvertTestSuite struct {
	suite.Suite
}

func TestConvertSuite(t *testing.T) {
	testSuite := &ConvertTestSuite{}
	suite.Run(t, testSuite)
}

func (suite *ConvertTestSuite) TestConvertComplexJsonConfig() {
	config, err := NewConfigFromJson(testdata.ValidJsonConfig2)
	assert.Nil(suite.T(), err)

	migration, err := NewMigrationFromConfig("", config)
	assert.Nil(suite.T(), err)

	_, err = migration.ConvertToYaml()
	assert.Nil(suite.T(), err)

}
