package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "5 + (3 * (2 - 4))", res)
	}
}

func Test2Operands(t *testing.T) {
	res, err := PostfixToInfix("3 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "5 + 3", res)
	}
}

func Test0Operands(t *testing.T) {
	_, err := PostfixToInfix("")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}

func Test1Operand(t *testing.T) {
	_, err := PostfixToInfix("+")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}

func TestThreeOperations(t *testing.T) {
	res, err := PostfixToInfix("119201979 1001818 - 4 / 97532367 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "97532367 + (4 / (1001818 - 119201979))", res)
	}
}

func TestSevenOperations(t *testing.T) {
	res, err := PostfixToInfix("1 2 + 3 - 4 * 5 / 6 + 7 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "7 - (6 + (5 / (4 * (3 - (2 + 1)))))", res)
	}
}

func TestSpaceOnly(t *testing.T) {
	_, err := PostfixToInfix(" ")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}

func Test2Operand(t *testing.T) {
	_, err := PostfixToInfix("2 +")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}

func TestUnallowedChars(t *testing.T) {
	_, err := PostfixToInfix("5 2 , 7 =")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
	  assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}
