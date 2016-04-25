// +build integration

package gull

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntegrationConvertSuite struct {
	suite.Suite
}

func TestIntegrationConvertSuite(t *testing.T) {
	testSuite := &IntegrationConvertSuite{}
	suite.Run(t, testSuite)
}

func (suite *IntegrationMigrateSuite) TestConvertConfigWithEnvironmentsAtTheRoot() {
	transform, err := NewConvert("/tmp/test-env-root-json/")
	assert.Nil(suite.T(), err)

	err = transform.ConvertDirectory("testdata/test-env-root-json")

	assert.Nil(suite.T(), err)
}
