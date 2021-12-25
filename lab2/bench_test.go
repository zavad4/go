package lab2

import (
	"fmt"
	"math/rand"
	"testing"
)

var startStatement string = "1 2 +"
var cntRes string
var err error

func BenchmarkPostfixToInfix(b *testing.B) {
	const baseLen = 150
	iterationsNum := baseLen
	
	signs := [4]string{"+", "-", "*", "/"}
	numbers := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < 20; i++ {
		iterationsNum = iterationsNum * 5 / 4
		input := startStatement

		for j := 0; j < iterationsNum; j++ {
			number := numbers[rand.Intn(10)]
			sign := signs[rand.Intn(4)]
	
			input = input + " " + number + " " + sign
		}
		b.Run(fmt.Sprintf("len=%d", iterationsNum), func(b *testing.B) {
			cntRes, err = PostfixToInfix(input)
		})
	}
}