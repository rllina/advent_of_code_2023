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

func (suite *mainSuite) Test_gets_the_first_and_last_numbers_from_a_string() {
	numbers := getNumbersFromString("two1nine")

	require.Equal(suite.T(), 29, numbers)
}

func (suite *mainSuite) Test_gets_the_numbers_from_multiple_lines() {
	numbers := getNumbersFromLines("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")

	require.Equal(suite.T(), 281, numbers)
}
