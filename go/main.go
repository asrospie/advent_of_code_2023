package main

import (
    "fmt"
)

func generator(msg string) <-chan string {
    ch := make(chan string)
    go func() {
        for i := 0; ; i++ {
            ch <- fmt.Sprintf("%s %d", msg, i)
        }
    }()
    return ch
}

func main() {
    fmt.Println("Advent of Code 2023!")
    fmt.Println("To run my solutions, use the test.sh script.")
    fmt.Println("NOTE: Some tests may take a while to run, use test_example.sh to only run the examples.")

    ch := generator("Hello")
    for i := 0; i < 5; i++ {
        fmt.Println(<-ch)
    }
}
