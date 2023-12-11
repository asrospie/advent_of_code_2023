package day11

import (
    "fmt"
    utils "rospierski/aocgo/pkg/aocutils"
)

type Coord struct {
    X int
    Y int
}

func (c *Coord) ManhattanDistance(other Coord) int {
    x_dist := c.X - other.X
    y_dist := c.Y - other.Y
    if x_dist < 0 {
        x_dist = -x_dist
    }
    if y_dist < 0 {
        y_dist = -y_dist
    }
    return x_dist + y_dist
}

func (c *Coord) String() string {
    return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func (c *Coord) Equals(other Coord) bool {
    return c.X == other.X && c.Y == other.Y
}

func getPairs(grid *[][]rune) []Coord {
    var coords []Coord
    for y, row := range *grid {
        for x, col := range row {
            if col == '#' {
                coords = append(coords, Coord{x, y})
            }
        }
    }
    return coords
}

func getDistancesSum(coords []Coord) int {
    sum := 0
    for i := 0; i < len(coords); i++ {
        for j := i + 1; j < len(coords); j++ {
            sum += coords[i].ManhattanDistance(coords[j])
        }
    } 
    return sum
}

func Day11Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
    coords := getPairs(&grid)
    empty_cols, empty_rows := findEmptyColsRows(&grid)
    new_coords := applyMultiplier(coords, 2, empty_cols, empty_rows)
    for i, coord := range new_coords {
        fmt.Printf("%d :: %s\n", i + 1, coord.String())
    }

    return getDistancesSum(new_coords), nil
}

func applyMultiplier(coords []Coord, multiplier int, empty_cols []int, empty_rows []int) []Coord {
    new_coords := make([]Coord, len(coords))
    for i, coord := range coords {
        x := coord.X
        y := coord.Y

        x_offset := 0
        for _, col := range empty_cols {
            if x > col {
                x_offset++
            }
            if x <= col {
                break
            } 
        }

        y_offset := 0
        for _, row := range empty_rows {
            if y > row {
                y_offset++
            }
            if y <= row {
                break
            }
        }

        fmt.Printf("x_offset: %d, y_offset: %d\n", x_offset, y_offset)
        x_after := x + (multiplier - 1) * x_offset
        y_after := y + (multiplier - 1) * y_offset
        fmt.Printf("%s -> (%d, %d)\n", coord.String(), x_after, y_after)
        new_coords[i] = Coord{x_after, y_after}
    }
    return new_coords
}

func findEmptyColsRows(grid *[][]rune) ([]int, []int) {
    var empty_cols []int
    var empty_rows []int

    for y, row := range *grid {
        row_check := true
        for _, col := range row {
            if col != '.' {
                row_check = false 
                continue
            }
        }
        if row_check {
            empty_rows = append(empty_rows, y)
        }
    }

    // iterate over columns, then by row
    for x := 0; x < len((*grid)[0]); x++ {
        col_check := true
        for y := 0; y < len(*grid); y++ {
            if (*grid)[y][x] != '.' {
                col_check = false
                continue
            }
        }
        if col_check {
            empty_cols = append(empty_cols, x)
        }
    }
    return empty_cols, empty_rows
}

func Day11Part2(filename string, multiplier int) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
    
    coords := getPairs(&grid)
    empty_cols, empty_rows := findEmptyColsRows(&grid)
    new_coords := applyMultiplier(coords, multiplier, empty_cols, empty_rows)

    for _, n := range new_coords {
        fmt.Printf("%s\n", n.String())
    }

    sum := getDistancesSum(new_coords)
    fmt.Printf("Sum: %d\n", sum)

    return sum, nil
}
