package day6

import (
	"errors"
	utils "rospierski/aocgo/pkg/aocutils"
	"strconv"
	"strings"
    "fmt"
)

type BoatRace struct {
    time int;
    dist int;
}

func NewBoatRace(time int, dist int) BoatRace {
    return BoatRace{time, dist}
}

func (b *BoatRace) PossibleWins() int {
    sum := 0
    for i := 1; i < b.time; i++ {
        if i * (b.time - i) > b.dist {
            sum += 1
        }
    }
    return sum
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

    fmt.Println(possibleWins)

    result := utils.ReduceSlice(possibleWins, func(a int, b int) int {
        if a == 0 {
            a = 1
        }
        return a * b
    })
    return result, nil
}
