package main

import (
    "fmt"
)

func generator(msg string) chan string {
    ch := make(chan string)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- fmt.Sprintf("%s %d", msg, i)
        }
        close(ch)
    }()
    return ch
}

func main() {
    fmt.Println("Advent of Code 2023!")
    fmt.Println("To run my solutions, use the test.sh script.")
    fmt.Println("NOTE: Some tests may take a while to run, use test_example.sh to only run the examples.")
}
