package lab2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (handler *ComputeHandler) Compute() error {
	reader := bufio.NewReader(handler.Input)
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	result, err := PostfixToInfix(str)
	if err == nil {
		fmt.Fprintln(handler.Output, result)
	} else {
		fmt.Fprintln(handler.Output, err)
	}
	return err
}
