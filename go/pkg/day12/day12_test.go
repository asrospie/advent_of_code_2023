package day12

import (
    "testing"
)

func TestDay12Part1Example(t *testing.T) {
    num, err := Day12Part1("../../inputs/day_12_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 21
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay12Part1Input(t *testing.T) {
    num, err := Day12Part1("../../inputs/day_12_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 7922
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}
