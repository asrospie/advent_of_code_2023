package day7

import (
	"fmt"
	utils "rospierski/pkg/aocutils"
)

func Day7Part1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return -1, err
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	return -1, nil
}
