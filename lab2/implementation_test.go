package lab2

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2Operands(t *testing.T) {
	res, err := PostfixToInfix("3 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "5 + 3", res)
	}
}
