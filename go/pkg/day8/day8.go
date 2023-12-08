package day8

import (
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
	"strings"
    "regexp"
    "sync"
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
    r, err := regexp.Compile(`[A-Z0-9]+`)
    if err != nil {
        return nil, err
    }

    for _, line := range lines {
        matches := r.FindAllString(line, -1)
        node_map[matches[0]] = []string{matches[1], matches[2]}
    }
    return node_map, nil
}

func getStartEndNodes(node_map map[string][]string) ([]string, []string) {
    var starting_nodes []string
    var ending_nodes []string
    for k := range node_map {
        if k[2] == 'A' {
            starting_nodes = append(starting_nodes, k)
        }
        if k[2] == 'Z' {
            ending_nodes = append(ending_nodes, k)
        }
    }
    return starting_nodes, ending_nodes
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

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(ints []int) int {
    if len(ints) == 0 {
        return 0
    }
    result := ints[0]
    for i := 1; i < len(ints); i++ {
        result = result * ints[i] / gcd(result, ints[i])
    }
    return result
}

func Day8Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    dirs := getDirections(lines[0])
    fmt.Println(dirs)
    node_map, err := getNodeMap(lines[2:])
    if err != nil {
        return -1, err
    }
    start_nodes, _ := getStartEndNodes(node_map)


    final_counters := make([]int, len(start_nodes))
    var wg sync.WaitGroup
    for i := 0; i < len(start_nodes); i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            dir_counter := 0
            counter := 0
            cur := start_nodes[i]
            for cur[2] != 'Z' {
                cur = node_map[cur][dirs[dir_counter]] 
                dir_counter += 1
                if dir_counter == len(dirs) {
                    dir_counter = 0
                }
                counter += 1
            }
            final_counters[i] = counter 
        }(i)
    }
    wg.Wait()

    fmt.Println(final_counters)

    result := lcm(final_counters)
    return result, nil
}
