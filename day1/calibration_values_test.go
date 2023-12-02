package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindingNumberSimple(t *testing.T) {
	assert := assert.New(t)

	expected, err := findFirstNumber("1abc2")
	assert.Equal(12, expected)
	assert.Nil(err)
}

func TestFindingNumberHarder(t *testing.T) {
	assert := assert.New(t)

	expected, err := findFirstNumber("pqr3stu8vwx")
	assert.Equal(12, expected)
	assert.Nil(err)
}
