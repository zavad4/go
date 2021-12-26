package commands

import (
	"fmt"
	"testing"
	"strconv"

	"github.com/zavad4/go/lab4/engine"
)

var cmd engine.Command

func BenchmarkParse(b *testing.B) {
	for i := 1; i <= 20; i++ {
		b.Run(fmt.Sprintf("%d-length", i*1000000), func(b *testing.B) {
			for j:=1; j<=1000; j++{
				stringNum := strconv.Itoa(i * 1000000)
			cmd = Parse("printc " + stringNum + " A")
			}
			
		})
	}
}