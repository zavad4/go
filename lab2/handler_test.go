package lab2

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	lab2 "github.com/zavad4/go/tree/main/lab2"
)

func TestConsoleInputCorrect(t *testing.T) {
	stdout := os.Stdout
	read, write, _ := os.Pipe()
	os.Stdout = write
	handler := lab2.ComputeHandler{
		Input:  bytes.NewBufferString("3 5 +"),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	out, _ := ioutil.ReadAll(read)
	os.Stdout = stdout
	if assert.Nil(t, err) {
		assert.Equal(t, "3 + 5\n", string(out[:]))
	} else {
		t.Errorf("Wrong result")
	}
}

func TestConsoleWrongChars(t *testing.T) {
	stdout := os.Stdout
	_, write, _ := os.Pipe()
	os.Stdout = write
	handler := lab2.ComputeHandler{
		Input:  bytes.NewBufferString("aaa aa"),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	//out, _ := ioutil.ReadAll(read)
	os.Stdout = stdout
	if assert.NotNil(t, err) {
		assert.Equal(t, "Unable to convert", err.Error())
	} else {
		t.Errorf("Wrong result")
	}
}

func TestConsoleWrongNumberArgs(t *testing.T) {
	stdout := os.Stdout
	_, write, _ := os.Pipe()
	os.Stdout = write
	handler := lab2.ComputeHandler{
		Input:  bytes.NewBufferString("3 5 "),
		Output: write,
	}
	err := handler.Compute()
	write.Close()
	//out, _ := ioutil.ReadAll(read)
	os.Stdout = stdout
	if assert.NotNil(t, err) {
		assert.Equal(t, "Unable to convert", err.Error())
	} else {
		t.Errorf("Wrong result")
	}
}
