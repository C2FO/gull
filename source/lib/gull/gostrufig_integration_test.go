// +build integration

// In order for these tests to pass, etcd needs to be running

package gull

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/brockwood/gostrufig"
	"github.com/brockwood/gostrufig/driver/etcd"

	"github.com/c2fo/gull/source/lib/gull/testdata"
)

type IntegrationGostrufigSuite struct {
	suite.Suite
	Target MigrationTarget
}

func TestIntegrationGostrufigSuite(t *testing.T) {
	migrateSuite := &IntegrationGostrufigSuite{
		Target: NewEtcdMigrationTarget(testdata.ValidEtcdHostUrl, "gull", "default"),
	}
	suite.Run(t, migrateSuite)
	_ = os.RemoveAll(testdata.ConvertDestination1)
}

func (suite *IntegrationGostrufigSuite) TestGostrufigRetrieve() {
	transform, err := NewConvert(testdata.ConvertDestination2)
	assert.Nil(suite.T(), err)

	err = transform.ConvertDirectory(testdata.ConvertSource2)
	assert.Nil(suite.T(), err)

	up := NewUp(testdata.ConvertDestination2, suite.Target)
	up.Migrate()

	var config testdata.GostrufigTestConfig1
	etcdDriver := etcd.GetGostrufigDriver()
	etcdUrl := strings.Replace(testdata.ValidEtcdHostUrl, "/v2/keys", "", -1)
	gostrufig := gostrufig.GetGostrufig("gull", etcdUrl, etcdDriver)
	gostrufig.RetrieveConfig(&config)
	assert.Equal(suite.T(), "carroll", config.Lewis)
}
