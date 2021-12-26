package commands

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zavad4/go/lab4/engine"
)

var constInput = "printc"
var command engine.Command

func BenchmarkCount(b *testing.B) {
	length := 1
	for i := 0; i < 27; i++ {
		length = 2 * length
		input := constInput
		input += strings.Repeat("A", length)

		b.Run(fmt.Sprintf("len=%d", length), func(b *testing.B) {
			command = Parse(input)
		})
	}
}
