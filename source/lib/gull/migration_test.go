package gull

import (
	"testing"

	"github.com/c2fo/gull/source/lib/gull/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MigrationTestSuite struct {
	suite.Suite
}

func TestMigrationSuite(t *testing.T) {
	suite.Run(t, new(MigrationTestSuite))
}

func (suite *MigrationTestSuite) TestApplyMigration() {
	result, err := NewMigrationFromYaml(testdata.ValidYamlMigration1, true)

	assert.Nil(suite.T(), err)
	err = result.ApplyUp()
	assert.Nil(suite.T(), err)
}
