package day6

import (
	"errors"
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
	"strconv"
	"strings"
)

type BoatRace struct {
    time int;
    dist int;
}

func NewBoatRace(time int, dist int) BoatRace {
    return BoatRace{time, dist}
}

func (b *BoatRace) PossibleWins() int {
    start := -1
    end := -1
    for i := 1; i < b.time; i++ {
        if i * (b.time - i) > b.dist {
            start = i
            break
        }
    }
    for i := b.time - 1; i > 0; i-- {
        if i * (b.time - i) > b.dist {
            end = i
            break
        }
    }
    return end - start + 1
}

func (b *BoatRace) String() string {
    return fmt.Sprintf("BoatRace{time: %d, dist: %d}", b.time, b.dist)
}

func getNums(line string) ([]int, error) {
    words := strings.Fields(line)
    times := utils.Mapper(words[1:], func(word string) int {
        num, err := strconv.Atoi(word)
        if err != nil {
            return -1
        }
        return num
    })
    if utils.SliceContains(times, -1) {
        return nil, errors.New("Could not convert all words to ints")
    }

    return times, nil
}

func getBoatRace(lines []string) (BoatRace, error) {
    times := strings.Fields(lines[0])[1:]
    dists := strings.Fields(lines[1])[1:]
    t, err := strconv.Atoi(strings.Join(times, ""))
    d, err := strconv.Atoi(strings.Join(dists, ""))

    return NewBoatRace(t, d), err
}

func getBoatRaces(times []int, dists []int) []BoatRace {
    boatRaces := make([]BoatRace, len(times))
    for i := 0; i < len(times); i++ {
        boatRaces[i] = NewBoatRace(times[i], dists[i])
    }
    return boatRaces
}

func Day6Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }
    times, err := getNums(lines[0])
    dists, err := getNums(lines[1])
    if err != nil {
        return -1, err
    }
    boatRaces := getBoatRaces(times, dists)
    possibleWins := utils.Mapper(boatRaces, func(boatRace BoatRace) int {
        return boatRace.PossibleWins()
    })

    result := utils.ReduceSlice(possibleWins, func(a int, b int) int {
        if a == 0 {
            a = 1
        }
        return a * b
    })
    return result, nil
}

func Day6Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    b, err := getBoatRace(lines)
    if err != nil {
        return -1, err
    }
    return b.PossibleWins(), nil
}
