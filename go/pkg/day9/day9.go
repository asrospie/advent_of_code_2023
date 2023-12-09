package day9

import (
    "strconv"
    "strings"
    utils "rospierski/aocgo/pkg/aocutils"
)

func strsToInts(lines []string) [][]int {
    return utils.Mapper(lines, func(line string) []int {
        ss := strings.Split(line, " ")
        xs := make([]int, len(ss))
        for i, s := range ss {
            num, err := strconv.Atoi(s)
            if err != nil {
                panic(err)
            }
            xs[i] = num
        }
        return xs
    })
}

func isAllZero(xs []int) bool {
    ret := true
    for _, x := range xs {
        ret = ret && (x == 0)
    }
    return ret
}

func lastEl[T any](xs []T) T {
    return xs[len(xs) - 1]
}

func getDiffs(below []int) []int {
    len_below := len(below)
    diffs := make([]int, len_below - 1)
    last := below[0]
    for i := 1; i < len_below; i++ {
        diffs[i-1] = below[i] - last
        last = below[i]
    }
    return diffs
}

func predictFutureHelper(above []int, below []int) int {
    if isAllZero(below) {
        return lastEl(above) + lastEl(below)
    }
    diffs := getDiffs(below)
    return predictFutureHelper(below, diffs) + lastEl(above)
}

func predictFuture(xs []int) int {
    return predictFutureHelper([]int{0}, xs)
}

func predictPastHelper(above []int, below[] int) int {
    if isAllZero(below) {
        d := above[0] - below[0]
        return d
    }
    diffs := getDiffs(below)
    return above[0] - predictPastHelper(below, diffs)
}

func predictPast(xs []int) int {
    diffs := getDiffs(xs)
    return predictPastHelper(xs, diffs)
}

func Day9Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    histories := strsToInts(lines)

    res := utils.SumSlice(utils.Mapper(histories, func(xs []int) int {
        return predictFuture(xs)
    }))

    return res, nil
}

func Day9Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    histories := strsToInts(lines)

    res := utils.SumSlice(utils.Mapper(histories, func(xs []int) int {
        return predictPast(xs)
    }))

    return res, nil
}
