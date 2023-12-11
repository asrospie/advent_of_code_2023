package day11

import (
    "testing"
)

func TestDay11Part1Example(t *testing.T) {
    num, err := Day11Part1("../../inputs/day_11_example.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 374
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay11Part1Input(t *testing.T) {
    num, err := Day11Part1("../../inputs/day_11_input.txt")
    if err != nil {
        t.Error(err)
    }
    expected := 10228230
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay11Part2Example1(t *testing.T) {
    num, err := Day11Part2("../../inputs/day_11_example.txt", 10)
    if err != nil {
        t.Error(err)
    }
    expected := 1030
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay11Part2Example2(t *testing.T) {
    num, err := Day11Part2("../../inputs/day_11_example.txt", 100)
    if err != nil {
        t.Error(err)
    }
    expected := 8410
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay11Part2Example3(t *testing.T) {
    num, err := Day11Part2("../../inputs/day_11_input.txt", 2)
    if err != nil {
        t.Error(err)
    }
    expected := 10228230
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}

func TestDay11Part2Input(t *testing.T) {
    num, err := Day11Part2("../../inputs/day_11_input.txt", 1_000_000)
    if err != nil {
        t.Error(err)
    }
    expected := 447073334102
    if num != expected {
        t.Errorf("Expected 0, got %d", num)
    }
}
