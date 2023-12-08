package day8

import (
	// "errors"
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
	// "strconv"
	"strings"
    "regexp"
)

func getDirections(line string) []int {
    dir_str := strings.Split(line, "")

    return utils.Mapper(dir_str, func(s string) int {
        if s == "L" {
            return 0
        }
        return 1
    })
}

func getNodeMap(lines []string) (map[string][]string, error) {
    node_map := make(map[string][]string)
    r, err := regexp.Compile(`[A-Z]+`)
    if err != nil {
        return nil, err
    }

    for _, line := range lines {
        matches := r.FindAllString(line, -1)
        node_map[matches[0]] = []string{matches[1], matches[2]}
    }
    return node_map, nil
}

func Day8Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename) 
    if err != nil {
        return -1, err
    }
    
    dirs := getDirections(lines[0])

    node_map, err := getNodeMap(lines[2:])
    if err != nil {
        return -1, err
    }

    start := "AAA"
    end := "ZZZ"
    cur := start
    
    dir_counter := 0
    counter := 0
    for cur != end {
        cur = node_map[cur][dirs[dir_counter]] 
        dir_counter += 1
        if dir_counter == len(dirs) {
            dir_counter = 0
        }
        counter += 1
    }

    return counter, nil
}
