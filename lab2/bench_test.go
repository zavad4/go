package lab2

import (
	"fmt"
	"testing"
)

var startStatement string = "2 1 +"
var cntRes string
var err error

func BenchmarkPostfixToInfix(b *testing.B) {
	const baseLen = 2500

	for i := 0; i < 25; i++ {
		input := startStatement
		iterationsNum := baseLen * (i + 1)

		for j := 0; j < iterationsNum; j++ {
			input = "+ " + input
			input = input + " "
			input = input + startStatement
		}

		b.Run(fmt.Sprintf("len=%d", iterationsNum), func(b *testing.B) {
			cntRes, err = PostfixToInfix(input)
		})
	}
}