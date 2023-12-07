package day7

import (
	"testing"
)

func TestDay7Part1Example(t *testing.T) {
	num, err := Day7Part1("../../inputs/day_7_example.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 6440
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}

func TestDay7Part1Input(t *testing.T) {
	num, err := Day7Part1("../../inputs/day_7_input.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 251058093
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}

func TestDay7Part2Example(t *testing.T) {
	num, err := Day7Part2("../../inputs/day_7_example.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 5905
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}


func TestDay7Part2Input(t *testing.T) {
	num, err := Day7Part2("../../inputs/day_7_input.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 251058093
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}
