package main

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mainSuite struct {
	suite.Suite
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(mainSuite))
}

func (suite *mainSuite) SetupTest() {}

func (suite *mainSuite) Test_does_stuff() {
	stuff := doStuff()

	require.Equal(suite.T(), stuff, "Sup")
}
