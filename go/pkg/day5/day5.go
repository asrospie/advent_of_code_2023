package day5

import (
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
	"strconv"
	"strings"
    "slices"
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

func getSeedRanges(line string) []chan int {
    seed_str := strings.Split(line, " ")[1:]
    var seeds []chan int
    for i := 0; i < len(seed_str); i += 2 {
        start, err := strconv.Atoi(seed_str[i])
        rg, err := strconv.Atoi(seed_str[i+1])
        if err != nil {
            panic(err)
        }

        ch := make(chan int)
        go func() {
            for j := start; j < start + rg; j++ {
                ch <- j
            }
            close(ch)
        }()
        seeds = append(seeds, ch)
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

    const MAX_INT = int(^uint(0) >> 1)
    minLocation := MAX_INT
    for _, seedGen := range seeds {
        for seed := range seedGen {
            l := seed
            for _, a := range almanac {
                l = a.GetDestination(l)
            }
            if l < minLocation {
                minLocation = l
            }
        } 
    }

    return minLocation, nil
}


func getAlmanacReverse(lines []string) ([]Map, error) {
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

        srcStart, err := strconv.Atoi(strSplit[0])
        destStart, err := strconv.Atoi(strSplit[1])
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
    slices.Reverse(almanac)
    return almanac, nil
}

type Seed struct {
    src_start int
    spread int
}

func (s *Seed) IsInSeed(i int) bool {
    return i >= s.src_start && i < s.src_start + s.spread
}

func getSeedRangesReverse(line string) []Seed {
    seed_str := strings.Split(line, " ")[1:]
    var seeds []Seed
    for i := 0; i < len(seed_str); i += 2 {
        start, err := strconv.Atoi(seed_str[i])
        rg, err := strconv.Atoi(seed_str[i+1])
        if err != nil {
            panic(err)
        }

        seeds = append(seeds, Seed{
            src_start: start,
            spread: rg,
        })
    }
    return seeds
}

func Day5Part2Reverse(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    lines = utils.Filter(lines, func(s string) bool { return len(s) > 0 })
    seeds := getSeedRangesReverse(lines[0])
    almanac, err := getAlmanacReverse(lines[1:])
    if err != nil {
        return -1, err
    }

    l := 0
    for {
        seed := l
        for _, a := range almanac {
            seed = a.GetDestination(seed)
        }
        for _, s := range seeds {
            if s.IsInSeed(seed) {
                return l, nil
            }
        }
        l += 1
    }
}
