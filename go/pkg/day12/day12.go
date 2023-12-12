package day12

import (
    "fmt"
    utils "rospierski/aocgo/pkg/aocutils"
    "strings"
    "strconv"
    "regexp"
)

func parseInput(lines []string) ([]string, [][]int) {
    rows := make([]string, len(lines))
    records := make([][]int, len(lines))
    for i, line := range lines {
        line_split := strings.Split(line, " ")
        rows[i] = line_split[0]

        num_str_split := strings.Split(line_split[1], ",")
        records[i] = utils.Mapper(num_str_split, func(s string) int {
            num, err := strconv.Atoi(s)
            if err != nil {
               panic(err)
            }
            return num
        })
    }
    return rows, records
}

func findPermutations(line string, record []int) int {
    unknown_len := strings.Count(line, "?") 
    max_bin_num := (1 << unknown_len) - 1

    binary_str := strings.ReplaceAll(line, "#", "0")
    binary_str = strings.ReplaceAll(binary_str, ".", "1")

    // build regex
    re_str := "1*"
    for i, num := range record {
        re_str += fmt.Sprintf("0{%d}", num)
        if i < len(record) - 1 {
            re_str += "1+"
            continue
        }
        re_str += "1*"
    }

    re := regexp.MustCompile(re_str)

    sum := 0
    for i := 0; i <= max_bin_num; i++ {
        temp := binary_str
        test := fmt.Sprintf("%b", i)
        for j := len(test); j < unknown_len; j++ {
            test = "0" + test
        }
        for _, c := range strings.Split(test, "") {
            temp = strings.Replace(temp, "?", c, 1) 
        }
        // check match
        match := re.FindString(temp)
        if match != "" && len(match) == len(temp) {
            sum++
        }
    }

    return sum
}

func Day12Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    rows, records := parseInput(lines)

    sum := 0
    for i, row := range rows {
        sum += findPermutations(row, records[i])
    }
    return sum, nil
}

func Day12Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    rows, records := parseInput(lines)


    return -1, nil
}
