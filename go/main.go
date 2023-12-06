package main

import (
    "fmt"
    "rospierski/aocgo/pkg/day6"
)

func main() {
    fmt.Println("Advent of Code 2023!")
    fmt.Println("To run my solutions, use the test.sh script.")
    fmt.Println("NOTE: Some tests may take a while to run, use test_example.sh to only run the examples.")

    result, err := day6.Day6Part1("./inputs/day_6_input.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("\nDay 6 Part 1: %d\n", result)
}
