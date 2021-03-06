package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((4 - 2) * 3) + 5)", res)
	}
}

func Test2Operands(t *testing.T) {
	res, err := PostfixToInfix("3 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "3 + 5", res)
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
		assert.Equal(t, "(((119201979 - 1001818) / 4) + 97532367)", res)
	}
}

func TestSevenOperations(t *testing.T) {
	res, err := PostfixToInfix("1 2 + 3 - 4 * 5 / 6 + 7 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "((((((1 + 2) - 3) * 4) / 5) + 6) - 7)", res)
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

func TestPowThreeOperations(t *testing.T) {
	res, err := PostfixToInfix("12 4 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, "12 ^ 4", res)
	}
}

func TestPowSevenOperations(t *testing.T) {
	res, err := PostfixToInfix("34 2 + 4 ^ 8 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((34 + 2) ^ 4) - 8)", res)
	}
}

func TestLetters(t *testing.T) {
	_, err := PostfixToInfix("3 b +")
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Unable to convert"), err)
	}
}
