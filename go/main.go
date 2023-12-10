package main

import (
    "fmt"
    "rospierski/aocgo/pkg/day10"
)

func main() {
    fmt.Println("Advent of Code 2023!")
    fmt.Println("To run my solutions, use the test.sh script.")
    fmt.Println("NOTE: Some tests may take a while to run, use test_example.sh to only run the examples.")

    result, err := day10.Day10Part1("./inputs/day_10_input.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}
