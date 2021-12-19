package commands

import (
	"fmt"

	"github.com/zavad4/go/lab4/engine"
)

type print struct {
	Arg string
}

func (p *print) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
