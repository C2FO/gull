// +build integration

// In order for these tests to run, etcd must be running

package gull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/c2fo/gull/source/lib/gull/testdata"
)

type IntegrationMigrateSuite struct {
	suite.Suite
	Target MigrationTarget
}

func TestIntegrationMigrateSuite(t *testing.T) {
	migrateSuite := &IntegrationMigrateSuite{
		Target: NewEtcdMigrationTarget(testdata.ValidEtcdHostUrl, "gull", "default"),
	}
	suite.Run(t, migrateSuite)
	_ = os.RemoveAll(testdata.ConvertDestination1)
}

func (suite *IntegrationMigrateSuite) TestLoadConfigIntoEtcd() {
	config, err := NewConfigFromJson(testdata.ValidJsonConfig1)
	assert.Nil(suite.T(), err)

	migration, err := NewMigrationFromConfig("", config)
	assert.Nil(suite.T(), err)

	err = migration.Content.Apply(suite.Target)
	assert.Nil(suite.T(), err)
}

func (suite *IntegrationMigrateSuite) TestMigrationStateStorageAndRetrieval() {
	transform, err := NewConvert(testdata.ConvertDestination1)
	assert.Nil(suite.T(), err)

	err = transform.ConvertDirectory(testdata.ConvertSource1)
	assert.Nil(suite.T(), err)

	up := NewUp(testdata.ConvertDestination1, suite.Target)
	err = up.Migrate()
	assert.Nil(suite.T(), err)

	state, err := suite.Target.GetStatus()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), state.Migrations)
	assert.Equal(suite.T(), up.Migrations.Count(), state.Migrations.Count())

	first, err := state.Migrations.First()
	assert.Nil(suite.T(), err)
	var leaf ConfigLeaf
	for _, entry := range first.Content.Entries {
		if entry.Path == "/default/alice" {
			leaf = entry
		}
	}
	assert.NotNil(suite.T(), leaf)
}

func (suite *IntegrationMigrateSuite) TestMigrateDown() {
	transform, err := NewConvert(testdata.ConvertDestination1)
	assert.Nil(suite.T(), err)

	err = transform.ConvertDirectory(testdata.ConvertSource1)
	assert.Nil(suite.T(), err)

	up := NewUp(testdata.ConvertDestination1, suite.Target)
	err = up.Migrate()
	assert.Nil(suite.T(), err)

	down := NewDown(suite.Target)
	err = down.Migrate()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), up.Migrations.Count()-1, down.Migrations.Count())

	upFirst, err := up.Migrations.First()
	assert.Nil(suite.T(), err)
	downFirst, err := down.Migrations.First()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), upFirst.Id, downFirst.Id)

	upLast, err := up.Migrations.Last()
	assert.Nil(suite.T(), err)
	downLast, err := down.Migrations.Last()
	assert.Nil(suite.T(), err)
	assert.NotEqual(suite.T(), upLast.Id, downLast.Id)
}
