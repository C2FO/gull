package gull

import (
	"testing"

	"github.com/c2fo/gull/source/lib/gull/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MigrationsTestSuite struct {
	suite.Suite
	Target MigrationTarget
}

func TestMigrationsSuite(t *testing.T) {
	mtSuite := &MigrationsTestSuite{
		Target: NewMockMigrationTarget("gull", "default", NewNullLogger()),
	}
	suite.Run(t, mtSuite)
}

func (suite *MigrationsTestSuite) TestPreviousMigrationLinks() {
	migrations := NewMigrations()
	migrationCount := 3
	for ii := 0; ii < migrationCount; ii++ {
		migration, err := NewMigrationFromGull("", testdata.ValidYamlMigration1)
		assert.Nil(suite.T(), err)

		err = migrations.Add(migration)
		assert.Nil(suite.T(), err)
	}

	last, err := migrations.Last()
	assert.Nil(suite.T(), err)

	first, err := migrations.First()
	assert.Nil(suite.T(), err)

	current := last
	for current.PreviousId != "" {
		current, err = migrations.Get(current.PreviousId)
		assert.Nil(suite.T(), err)
	}
	assert.Equal(suite.T(), first.Id, current.Id)
}

func (suite *MigrationsTestSuite) TestNextMigrationLinks() {
	migrations := NewMigrations()
	migrationCount := 3
	for ii := 0; ii < migrationCount; ii++ {
		migration, err := NewMigrationFromGull("", testdata.ValidYamlMigration1)
		assert.Nil(suite.T(), err)

		err = migrations.Add(migration)
		assert.Nil(suite.T(), err)
	}

	last, err := migrations.Last()
	assert.Nil(suite.T(), err)

	first, err := migrations.First()
	assert.Nil(suite.T(), err)

	current := first
	for current.NextId != "" {
		current, err = migrations.Get(current.NextId)
		assert.Nil(suite.T(), err)
	}
	assert.Equal(suite.T(), last.Id, current.Id)
}

func (suite *MigrationsTestSuite) TestApplyMultipleMigrations() {
	migrations := NewMigrations()
	migrationCount := 3
	for ii := 0; ii < migrationCount; ii++ {
		migration, err := NewMigrationFromGull("", testdata.ValidYamlMigration1)
		assert.Nil(suite.T(), err)
		err = migrations.Add(migration)
		assert.Nil(suite.T(), err)
	}

	err := migrations.Apply(suite.Target)
	assert.Nil(suite.T(), err)
}
