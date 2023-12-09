package day9

import (
    "testing"
)

func TestDay9Part1Example(t *testing.T) {
    num, err := Day9Part1("../../inputs/day_9_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 114
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay9Part1Input(t *testing.T) {
    num, err := Day9Part1("../../inputs/day_9_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 1853145119
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay9Part2Example(t *testing.T) {
    num, err := Day9Part2("../../inputs/day_9_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 2
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay9Part2Input(t *testing.T) {
    num, err := Day9Part2("../../inputs/day_9_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 923
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}
