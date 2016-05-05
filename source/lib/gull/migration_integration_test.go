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
		Target: NewEtcdMigrationTarget(testdata.ValidEtcdHostUrl, "default"),
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
	up.Migrate()

	state, err := suite.Target.GetStatus()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), state.Migrations)
	assert.Equal(suite.T(), up.Migrations.Count(), state.Migrations.Count())

	first, err := state.Migrations.First()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "/default/alice", first.Content.Entries[0].Path)
}
