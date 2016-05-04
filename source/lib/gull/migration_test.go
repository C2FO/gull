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
	mtSuite := new(MigrationTestSuite)
	mtSuite.Target = testdata.NewMockMigrationTarget("default")
	suite.Run(t, mtSuite)
}

func (suite *MigrationTestSuite) TestApplyYamlBasedMigration() {
	result, err := NewMigrationFromGull("", testdata.ValidYamlMigration1)
	assert.Nil(suite.T(), err)

	err = result.Content.Apply(suite.Target)
	assert.Nil(suite.T(), err)
}

func (suite *MigrationTestSuite) TestApplyConfigBasedMigration() {
	config, err := NewConfigFromJson(testdata.ValidJsonConfig1)
	assert.Nil(suite.T(), err)

	result, err := NewMigrationFromConfig("", config)
	assert.Nil(suite.T(), err)

	err = result.Content.Apply(suite.Target)
	assert.Nil(suite.T(), err)
}
