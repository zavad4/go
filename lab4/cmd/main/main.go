package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/zavad4/go/lab4/commands"
	"github.com/zavad4/go/lab4/engine"
)

func main() {
	path := "./input.txt"
	file := flag.String("f", path, "Input commands")
	flag.Parse()

	input, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Impossible to open file:  %s", err)
		return
	}
	defer input.Close()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := commands.Parse(commandLine)
		eventLoop.Post(cmd)
	}

	eventLoop.AwaitFinish()
}
