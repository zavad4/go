package main

import "fmt"

func main() {
	startSignal := make(chan int)
	go func() {
		<-startSignal
		fmt.Println("routine ended")
	}()
	fmt.Println("main started")
}
