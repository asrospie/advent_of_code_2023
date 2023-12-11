package day10

import (
    "testing"
)

func TestDay10Part1Example1(t *testing.T) {
    num, err := Day10Part1("../../inputs/day_10_example_1.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 4
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay10Part1Example2(t *testing.T) {
    num, err := Day10Part1("../../inputs/day_10_example_2.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 8
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay10Part1Input(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping test in short mode.")
    }
    num, err := Day10Part1("../../inputs/day_10_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 6768
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay10Part2Example1(t *testing.T) {
    num, err := Day10Part2("../../inputs/day_10_example_5.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 10
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay10Part2Input(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping test in short mode.")
    }
    num, err := Day10Part2("../../inputs/day_10_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 351
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}
