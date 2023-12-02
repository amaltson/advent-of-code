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

func TestAddingSimpleExample(t *testing.T) {
	assert := assert.New(t)

	simpleExample := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	result, err := addUpTheNumbers(simpleExample)
	assert.Equal(142, result)
	assert.Nil(err)
}

func TestAddingSimpleExampleWithEmpties(t *testing.T) {
	assert := assert.New(t)

	simpleExample := []string{"1abc2", "", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet", ""}
	result, err := addUpTheNumbers(simpleExample)
	assert.Equal(142, result)
	assert.Nil(err)
}

func TestConvertingWordNumberSimple(t *testing.T) {
	assert := assert.New(t)

	result := convertSpelledOutNumbers("two")
	assert.Equal("2", result)
}

func TestConvertingWordNumbers(t *testing.T) {
	assert := assert.New(t)

	result := convertSpelledOutNumbers("two1ninetwokkerca13numtwo")
	assert.Equal("2192kkerca13num2", result)
}

func TestConvertingOverlappingWordNumbers(t *testing.T) {
	assert := assert.New(t)

	result := convertSpelledOutNumbers("eightwothree")
	assert.Equal("8wo3", result)
}

func TestAddingWithWrittenDigits(t *testing.T) {
	assert := assert.New(t)

	writtenDigits := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four",
		"4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	result, err := addUpTheNumbers(writtenDigits)
	assert.Equal(281, result)
	assert.Nil(err)
}

func TestFindNumberOfStart(t *testing.T) {
	assert := assert.New(t)

	result := findNumberStart("eightwo5three")
	assert.Equal("8", result)
}

func TestFindNumberOfEnd(t *testing.T) {
	assert := assert.New(t)

	result := findNumberEnd("eightwo5three")
	assert.Equal("3", result)
}
