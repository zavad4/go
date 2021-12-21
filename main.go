package main

import "fmt"

func main() {
    data := make(chan int)
    go func() {
        for i := 10; i >= 0; i-- {
            data <- i
        }
    }()
    for i := 0; i < 5; i++ {
        fmt.Println(<-data)
    }
}