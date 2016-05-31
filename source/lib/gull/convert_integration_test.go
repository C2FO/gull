// +build integration

package gull

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/c2fo/gull/source/lib/gull/testdata"
)

type IntegrationConvertSuite struct {
	suite.Suite
}

func TestIntegrationConvertSuite(t *testing.T) {
	testSuite := &IntegrationConvertSuite{}
	suite.Run(t, testSuite)

	_ = os.RemoveAll(testdata.ConvertDestination1)
}

func (suite *IntegrationConvertSuite) TestConvertConfigWithEnvironmentsAtTheRoot() {
	transform, err := NewConvert(testdata.ConvertDestination1, false, false)
	assert.Nil(suite.T(), err)

	err = transform.ConvertDirectory(testdata.ConvertSource1)

	assert.Nil(suite.T(), err)
}
