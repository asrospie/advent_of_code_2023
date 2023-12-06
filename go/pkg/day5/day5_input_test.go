package day5

import (
    "testing"
)

func TestDay5Part1Input(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping test in short mode.")
    }
    num, err := Day5Part1("../../inputs/day_5_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 26273516
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay5Part2Input(t *testing.T) {
    t.Skip("Solution takes too long to run. See next test for faster solution.")
    num, err := Day5Part2("../../inputs/day_5_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 34039469
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}

func TestDay5Part2InputReverse(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping test in short mode.")
    }
    num, err := Day5Part2Reverse("../../inputs/day_5_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 34039469
    if num != expected {
        t.Errorf("Expected %d, got %d", expected, num)
    }
}
