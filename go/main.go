package main

import (
    "fmt"
    "rospierski/aocgo/pkg/day8"
)

func main() {
    fmt.Println("Advent of Code 2023!")
    fmt.Println("To run my solutions, use the test.sh script.")
    fmt.Println("NOTE: Some tests may take a while to run, use test_example.sh to only run the examples.")

    result, err := day8.Day8Part2("./inputs/day_8_input.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}
