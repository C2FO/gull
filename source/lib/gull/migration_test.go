package gull

import (
	"testing"

	"github.com/c2fo/gull/source/lib/gull/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MigrationTestSuite struct {
	suite.Suite
	Target MigrationTarget
}

func TestMigrationSuite(t *testing.T) {
	suite.Run(t, new(MigrationTestSuite))
}

func (suite *MigrationTestSuite) TestDryApplyYamlBasedMigration() {
	result, err := NewMigrationFromYaml(testdata.ValidYamlMigration1)
	assert.Nil(suite.T(), err)

	err = result.Up.Apply(true, suite.Target)
	assert.Nil(suite.T(), err)
}

func (suite *MigrationTestSuite) TestDryApplyConfigBasedMigration() {
	config, err := NewConfigFromJson(testdata.ValidJsonConfig1)
	assert.Nil(suite.T(), err)

	result, err := NewMigrationFromConfig(config, nil)
	assert.Nil(suite.T(), err)

	err = result.Up.Apply(true, suite.Target)
	assert.Nil(suite.T(), err)
}
