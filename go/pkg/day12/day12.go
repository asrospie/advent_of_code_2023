package day12

import (
	"fmt"
	"regexp"
	utils "rospierski/aocgo/pkg/aocutils"
	"strconv"
	"strings"
	"sync"
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

    sums := make([]int, max_bin_num + 1)
    wg := sync.WaitGroup{}
    for i := 0; i <= max_bin_num; i++ {
        wg.Add(1)
        go func(i int, max_bin_num int, unknown_len int, binary_str string, re *regexp.Regexp) {
            defer wg.Done()
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
                sums[i]++
            }
        }(i, max_bin_num, unknown_len, binary_str, re)
    }
    wg.Wait()
    sum = utils.SumSlice(sums)
    fmt.Println(sums)
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

func findPermutationsRec(line string, record []int, cache *map[string]int) int {
    if len(line) == 0 {
        if len(record) == 0 {
            return 1
        }
        return 0
    }

    if len(record) == 0 {
        if strings.Contains(line, "#") {
            return 0
        }
        return 1
    }

    key := fmt.Sprintf("%s:%v", line, record)
    if v, ok := (*cache)[key]; ok {
        return v
    }

    total := 0

    cs := []rune(line)
    c0 := cs[0]
    if c0 == '.' || c0 == '?' {
        total += findPermutationsRec(string(cs[1:]), record, cache)
    }

    if c0 == '#' || c0 == '?' {
        block_less_than_len := record[0] <= len(cs)
        block_has_no_dot := !strings.Contains(string(cs[:record[0]]), ".")
        var tester bool
        if record[0] >= len(cs) {
            tester = false
        } else {
            tester = cs[record[0]] != '#'
        }
        block_has_no_hash := len(cs) == record[0] || tester 
        if block_less_than_len && block_has_no_dot && block_has_no_hash {
            var next_cfg string
            if record[0] >= len(cs) {
                next_cfg = ""
            } else {
                next_cfg = string(cs[(record[0] + 1):])
            }
            total += findPermutationsRec(next_cfg, record[1:], cache)
        }
    }

    (*cache)[key] = total
    return total
}

func Day12Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    rows, records := parseInput(lines)
    for i := 0; i < len(rows); i++ {
        temp_row := rows[i]
        temp_record := records[i]
        for j := 0; j < 4; j++ {
            rows[i] += "?" + temp_row
            records[i] = append(records[i], temp_record...)
        }
    } 
    
    cache := make(map[string]int)

    sum := 0
    for i, row := range rows {
        sum += findPermutationsRec(row, records[i], &cache)
    }
    fmt.Println(sum)

    return sum, nil
}
