package day6

import (
    "testing"
)

func TestDay6Part1Example(t *testing.T) {
    num, err := Day6Part1("../../inputs/day_6_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 288
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay6Part1Input(t *testing.T) {
    num, err := Day6Part1("../../inputs/day_6_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 1731600
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay6Part2Example(t *testing.T) {
    num, err := Day6Part2("../../inputs/day_6_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 71503
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay6Part2Input(t *testing.T) {
    num, err := Day6Part2("../../inputs/day_6_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 40087680
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}
