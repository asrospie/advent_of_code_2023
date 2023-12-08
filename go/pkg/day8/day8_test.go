package day8

import (
	"testing"
)

func TestDay8Part1Example(t *testing.T) {
	num, err := Day8Part1("../../inputs/day_8_example.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 2
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}

func TestDay8Part1Example2(t *testing.T) {
	num, err := Day8Part1("../../inputs/day_8_example_2.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 6
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}

func TestDay8Part1Input(t *testing.T) {
	num, err := Day8Part1("../../inputs/day_8_input.txt")
	if err != nil {
		t.Error(err)
	}
	expected := 251058093
	if num != expected {
		t.Errorf("Expected %d, got %d", expected, num)
	}
}

// func TestDay8Part2Example(t *testing.T) {
// 	num, err := Day8Part2("../../inputs/day_8_example.txt")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	expected := 5905
// 	if num != expected {
// 		t.Errorf("Expected %d, got %d", expected, num)
// 	}
// }
//
//
// func TestDay8Part2Input(t *testing.T) {
// 	num, err := Day8Part2("../../inputs/day_8_input.txt")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	expected := 249781879
// 	if num != expected {
// 		t.Errorf("Expected %d, got %d", expected, num)
// 	}
// }
