package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

