package lab2

import (
	"fmt"
	//"math/rand"
	"testing"
)

var startStatement string = "1 2 +"
var cntRes string
var err error

func BenchmarkPostfixToInfix(b *testing.B) {
	const baseLen = 13
	iterationsNum := baseLen
	input := "1 34 + 53 + 13 + 43 454 / - 2 4 * 3 / 4 3 2 5 - * 6 / + / +"

	for i := 0; i < 18; i++ {
		b.Run(fmt.Sprintf("length = %d", iterationsNum), func(b *testing.B) {
			cntRes, err = PostfixToInfix(input)
		})

		iterationsNum = iterationsNum*2 + 1
		input = input + input + " +"
	}
}
