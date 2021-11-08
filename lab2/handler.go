package lab2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	In  io.Reader
	Out io.Writer
}

func Compute(ch *ComputeHandler) error {
	reader := bufio.NewReader(ch.In)

	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")

	result, err := PostfixToInfix(text)
	if err != nil {
		return err
	} else {
		fmt.Fprintln(ch.Out, "postfix:", result)
	}
	return nil
}
