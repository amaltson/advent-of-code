package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindingNumberSimple(t *testing.T) {
	assert := assert.New(t)

	result, err := findFirstAndLastNumber("1abc2")
	assert.Equal(12, result)
	assert.Nil(err)
}

func TestFindingNumberHarder(t *testing.T) {
	assert := assert.New(t)

	result, err := findFirstAndLastNumber("pqr3stu8vwx")
	assert.Equal(38, result)
	assert.Nil(err)
}

func TestExtraNumbers(t *testing.T) {
	assert := assert.New(t)

	result, err := findFirstAndLastNumber("a1b2c3d4e5f")
	assert.Equal(15, result)
	assert.Nil(err)
}

func TestOnlyOneNumber(t *testing.T) {
	assert := assert.New(t)

	result, err := findFirstAndLastNumber("treb7uchet")
	assert.Equal(77, result)
	assert.Nil(err)
}
