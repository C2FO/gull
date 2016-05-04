// +build integration

// In order for these tests to run, etcd must be running

package gull

import (
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
}

func (suite *IntegrationMigrateSuite) TestLoadConfigIntoEtcd() {
	config, err := NewConfigFromJson(testdata.ValidJsonConfig1)
	assert.Nil(suite.T(), err)

	migration, err := NewMigrationFromConfig("", config)
	assert.Nil(suite.T(), err)

	err = migration.Content.Apply(suite.Target)
	assert.Nil(suite.T(), err)
}
