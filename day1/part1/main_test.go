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

func (suite *mainSuite) Test_gets_the_numbers_from_a_string() {
	numbers := getNumbersFromString("1abc2")

	require.Equal(suite.T(), numbers, 12)
}

func (suite *mainSuite) Test_gets_the_first_and_last_numbers_from_a_string() {
	numbers := getNumbersFromString("a1b2c3d4e5f")

	require.Equal(suite.T(), numbers, 15)
}

func (suite *mainSuite) Test_gets_the_numbers_from_multiple_lines() {
	numbers := getNumbersFromLines("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")

	require.Equal(suite.T(), numbers, 142)
}

func (suite *mainSuite) Test_gets_the_numbers_from_the_part_1_input() {
	main()
}
