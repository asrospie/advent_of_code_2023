package aocutils

import (
    "os"
    "bufio"
    "cmp"
)

func ReadFileLines(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if scanner.Err() != nil {
        return nil, scanner.Err()
    }
    return lines, nil
}

func Filter[T any](s []T, test func(T) bool) []T {
    var ret []T
    for _, v := range s {
        if test(v) {
            ret = append(ret, v)
        }
    }
    return ret
}

func Mapper[T any, U any](s []T, f func(T) U) []U {
    var ret []U
    for _, v := range s {
        ret = append(ret, f(v))
    }
    return ret
}

func MinSlice[T cmp.Ordered](s []T) T {
    return ReduceSlice(s, func(a, b T) T {
        if a < b {
            return a
        }
        return b
    })
}

func MaxSlice[T cmp.Ordered](s []T) T {
    return ReduceSlice(s, func(a, b T) T {
        if a > b {
            return a
        }
        return b
    })
}

func SliceContains[T comparable](s []T, e T) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func SumSlice[T int | float32 | float64](s []T) T {
    return ReduceSlice(s, func(a, b T) T { return a + b })
}

func ReduceSlice[T any](s []T, f func(T, T) T) T {
    sum := s[0]
    for i := 1; i < len(s); i++ {
        sum = f(sum, s[i])
    }
    return sum
}
