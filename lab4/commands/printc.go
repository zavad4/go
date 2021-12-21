package commands

import (
	"strings"

	"github.com/zavad4/go/lab4/engine"
)

type printc struct {
	C   int
	Arg string
}

func (p *printc) Execute(loop engine.Handler) {
	res := strings.Repeat(p.Arg, p.C)
	loop.Post(&print{Arg: res})
}
