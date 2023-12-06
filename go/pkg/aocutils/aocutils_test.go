package aocutils

import (
    "testing"
)

func TestReduce(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    expected := 15

    actual := ReduceSlice(input, func(a, b int) int { return a + b })
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }
}

func TestSum(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    expected := 15

    actual := SumSlice(input)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }
}

func TestSliceContains(t *testing.T) {
    input := []string{"Alec", "Rospierski", "software", "engineer"}

    expected := false
    actual := SliceContains(input, "computer")

    if actual != expected {
        t.Errorf("Expected %t, got %t", expected, actual)
    }

    expected = true
    actual = SliceContains(input, "engineer")
    if actual != expected {
        t.Errorf("Expected %t, got %t", expected, actual)
    }
}

func TestMaxSlice(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    expected := 5

    actual := MaxSlice(input)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }
}

func TestMinSlice(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    expected := 1

    actual := MinSlice(input)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }
}

