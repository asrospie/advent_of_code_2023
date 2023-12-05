package day5

import (
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
	"strconv"
	"strings"
)

type Map struct {
    spread []int
    src_start []int
    dest_start []int
}

func (m *Map) GetDestination(src int) int {
    for i, s := range m.src_start {
        if src >= s && src < s + m.spread[i] {
            adder := src - s
            return m.dest_start[i] + adder 
        }
    }
    return src
}

func (m *Map) ToString() string {
    str := fmt.Sprintf("Map %d\n", len(m.src_start))
    for i, s := range m.src_start {
        str += fmt.Sprintf("%d %d %d\n", m.dest_start[i], s, m.spread[i])
    }
    return str
}

func getSeeds(line string) []int {
    seed_str := strings.Split(line, " ")[1:]
    seeds := utils.Mapper(seed_str, func(s string) int {
        i, err := strconv.Atoi(s)
        if err != nil {
            panic(err)
        }
        return i
    }) 
    return seeds
}

func getAlmanac(lines []string) ([]Map, error) {
    var almanac []Map

    var srcStarts []int
    var destStarts []int
    var spreads []int

    for idx, line := range lines {
        strSplit := strings.Split(line, " ")
        if _, err := strconv.Atoi(strSplit[0]); err != nil {
            if idx == 0 { continue }
            almanac = append(almanac, Map{
                spread: spreads,
                src_start: srcStarts,
                dest_start: destStarts,
            })
            srcStarts = []int{}
            destStarts = []int{}
            spreads = []int{}
            continue
        }

        destStart, err := strconv.Atoi(strSplit[0])
        srcStart, err := strconv.Atoi(strSplit[1])
        spread, err := strconv.Atoi(strSplit[2])

        if err != nil {
            return nil, err
        }

        srcStarts = append(srcStarts, srcStart)
        destStarts = append(destStarts, destStart)
        spreads = append(spreads, spread)
    }
    almanac = append(almanac, Map{
        spread: spreads,
        src_start: srcStarts,
        dest_start: destStarts,
    })
    return almanac, nil
}

func Day5Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    lines = utils.Filter(lines, func(s string) bool { return len(s) > 0 })
    seeds := getSeeds(lines[0])
    almanac, err := getAlmanac(lines[1:])
    if err != nil {
        return -1, err
    }

    var locations []int
    for _, seed := range seeds {
        l := seed
        for _, a := range almanac {
            l = a.GetDestination(l)
        }
        locations = append(locations, l)
    }

    return utils.MinSlice(locations), nil
}

func getSeedRanges(line string) []int {
    seed_str := strings.Split(line, " ")[1:]
    var seeds []int
    for i := 0; i < len(seed_str); i += 2 {
        start, err := strconv.Atoi(seed_str[i])
        rg, err := strconv.Atoi(seed_str[i+1])
        if err != nil {
            panic(err)
        }
        for j := start; j < start + rg; j++ {
            seeds = append(seeds, j)
        }
    }
    return seeds
}

func Day5Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    lines = utils.Filter(lines, func(s string) bool { return len(s) > 0 })
    seeds := getSeedRanges(lines[0])
    almanac, err := getAlmanac(lines[1:])
    if err != nil {
        return -1, err
    }

    var locations []int
    for _, seed := range seeds {
        l := seed
        for _, a := range almanac {
            l = a.GetDestination(l)
        }
        locations = append(locations, l)
    }

    return utils.MinSlice(locations), nil
}
