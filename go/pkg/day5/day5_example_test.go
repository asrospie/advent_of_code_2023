package day5

import (
    "testing"
)

func TestDay5Part1Example(t *testing.T) {
    num, err := Day5Part1("../../inputs/day_5_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 35
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay5Part2Example(t *testing.T) {
    num, err := Day5Part2("../../inputs/day_5_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 46
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay5Part2ReverseExample(t *testing.T) {
    num, err := Day5Part2("../../inputs/day_5_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 46
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}
