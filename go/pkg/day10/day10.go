package day10

import (
    "reflect"
    "fmt"
    utils "rospierski/aocgo/pkg/aocutils"
)

const (
    VERTICAL = '|' 
    HORIZONTAL = '-'
    BOTTOM_LEFT_BEND = 'L'
    BOTTOM_RIGHT_BEND = 'J'
    TOP_RIGHT_BEND = '7'
    TOP_LEFT_BEND = 'F'
    GROUND = '.'
    STARTING = 'S'
)

type Vertex struct {
    Type rune
    IsStart bool
    X int
    Y int
    Edges []*Edge
}

func (v *Vertex) String() string {
    s := fmt.Sprintf("Vertex (%v, %v) %v\n", v.X, v.Y, typeToString(v.Type))
    for _, e := range v.Edges {
        s += fmt.Sprintf("\tConnects (%v, %v) %v\n", e.Vertex.X, e.Vertex.Y, typeToString(e.Vertex.Type))
    }
    return s
}

type Edge struct {
    Vertex *Vertex
}

type Graph struct {
    Vertices []*Vertex
}

func (g *Graph) AddVertex(x int, y int, v_type rune) {
    v := Vertex{
        Type: v_type,
        X: x,
        Y: y,
        Edges: make([]*Edge, 0),
        IsStart: v_type == STARTING,
    }
    g.Vertices = append(g.Vertices, &v)
}

func (v *Vertex) CanConnectSource(search_dir int) bool {
    switch v.Type {
    case VERTICAL:
        return search_dir == UP || search_dir == DOWN
    case HORIZONTAL:
        return search_dir == LEFT || search_dir == RIGHT
    case BOTTOM_LEFT_BEND:
        return search_dir == UP || search_dir == RIGHT
    case BOTTOM_RIGHT_BEND:
        return search_dir == UP || search_dir == LEFT
    case TOP_RIGHT_BEND:
        return search_dir == DOWN || search_dir == LEFT
    case TOP_LEFT_BEND:
        return search_dir == DOWN || search_dir == RIGHT
    case STARTING:
        return true
    default:
        return false
    }
}

func (v *Vertex) CanConnectDestination(search_dir int) bool {
    switch v.Type {
    case VERTICAL:
        return search_dir == UP || search_dir == DOWN
    case HORIZONTAL:
        return search_dir == LEFT || search_dir == RIGHT
    case BOTTOM_LEFT_BEND:
        return search_dir == DOWN || search_dir == LEFT
    case BOTTOM_RIGHT_BEND:
        return search_dir == DOWN || search_dir == RIGHT
    case TOP_RIGHT_BEND:
        return search_dir == UP || search_dir == RIGHT
    case TOP_LEFT_BEND:
        return search_dir == UP || search_dir == LEFT
    case STARTING:
        return true
    default:
        return false
    }
}

func dirToString(dir int) string {
    switch dir {
    case UP:
        return "UP"
    case RIGHT:
        return "RIGHT"
    case DOWN:
        return "DOWN"
    case LEFT:
        return "LEFT"
    default:
        return "unknown"
    }
}

func (g *Graph) AddEdge(v1 *Vertex, v2 *Vertex, search_dir int) {
    // check if can connect
    if !(v2.CanConnectDestination(search_dir) && v1.CanConnectSource(search_dir)) {
        return
    }
    // check if edge alread exists
    for _, e := range v1.Edges {
        if reflect.DeepEqual(e.Vertex, v2) {
            return
        }
    }
    e := Edge{
        Vertex: v2,
    }
    v1.Edges = append(v1.Edges, &e)
}

func (v *Vertex) IsEqual(v2 *Vertex) bool {
    return v.X == v2.X && v.Y == v2.Y
}

func (g *Graph) FindFurthestVertexDistance(v *Vertex) int {
    // find furthest vertex
    counter := 0
    current := v.Edges[0].Vertex
    prev := v
    start_v := g.GetStartingVertex()
    for reflect.DeepEqual(current, start_v) == false {
        next := current.Edges[0].Vertex
        next_is_prev := next.IsEqual(prev)
        if next_is_prev {
            next = current.Edges[1].Vertex
        }
        prev = current
        current = next
        counter++
    }
    
    if counter % 2 == 0 {
        return counter / 2
    }
    return counter / 2 + 1
}

func (g *Graph) GetVertex(x int, y int) *Vertex {
    for _, v := range g.Vertices {
        if v.X == x && v.Y == y {
            return v
        }
    }
    return nil
}

func (g *Graph) String() string {
    s := ""
    for _, v := range g.Vertices {
        s += fmt.Sprintf("%s\n", v.String())
    }
    return s
}

func buildGraph(grid [][]rune) *Graph {
    g := Graph{
        Vertices: make([]*Vertex, 0),
    }
    for y, row := range grid {
        for x, v := range row {
            if v == GROUND {
                continue
            }
            g.AddVertex(x, y, v)
        }
    }
    for y, row := range grid {
        for x, v := range row {
            if v == GROUND {
                continue
            }

            ver := g.GetVertex(x, y)
            if ver == nil {
                continue
            }
            if v2 := g.GetVertex(x+1, y); v2 != nil {
                g.AddEdge(ver, v2, RIGHT)
            }
            if v2 := g.GetVertex(x-1, y); v2 != nil {
                g.AddEdge(ver, v2, LEFT)
            }
            if v2 := g.GetVertex(x, y+1); v2 != nil {
                g.AddEdge(ver, v2, DOWN)
            }
            if v2 := g.GetVertex(x, y-1); v2 != nil {
                g.AddEdge(ver, v2, UP)
            }
        }
    }
    return &g
}

func (g *Graph) GetStartingVertex() *Vertex {
    for _, v := range g.Vertices {
        if v.IsStart {
            return v
        }
    }
    return nil
}

func typeToString(t rune) string {
    switch t {
    case VERTICAL:
        return "[|]"
    case HORIZONTAL:
        return "[-]"
    case BOTTOM_LEFT_BEND:
        return "[L]"
    case BOTTOM_RIGHT_BEND:
        return "[J]"
    case TOP_RIGHT_BEND:
        return "[7]"
    case TOP_LEFT_BEND:
        return "[F]"
    case GROUND:
        return "[.]"
    case STARTING:
        return "[S]"
    default:
        return "unknown"
    }
}

const (
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3
)



func getCharGrid(lines []string) [][]rune {
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
    return grid
}


func Day10Part1(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    grid := getCharGrid(lines)
    g := buildGraph(grid)
    v := g.GetStartingVertex()
    return g.FindFurthestVertexDistance(v), nil
}

func walkGraph(g *Graph, start *Vertex, grid *[][]rune) {
    // find furthest vertex
    current := start.Edges[0].Vertex
    prev := start
    start_v := g.GetStartingVertex()
    (*grid)[current.Y][current.X] = 'X'
    for reflect.DeepEqual(current, start_v) == false {
        next := current.Edges[0].Vertex
        next_is_prev := next.IsEqual(prev)
        if next_is_prev {
            next = current.Edges[1].Vertex
        }
        prev = current
        current = next
        switch current.Type {
        case VERTICAL:
            fallthrough
        case BOTTOM_RIGHT_BEND:
            fallthrough
        case BOTTOM_LEFT_BEND:
            (*grid)[current.Y][current.X] = 'W'
        default:
            (*grid)[current.Y][current.X] = 'X'
        }
    }
    for y, row := range *grid {
        for x, v := range row {
            if v == 'X' || v == 'W' {
                continue
            }
            (*grid)[y][x] = '.'
        }
    }
}

func scanForX(x, y int, grid *[][]rune) {
    min_x := 0 
    max_x := len((*grid)[0])
    min_y := 0
    max_y := len(*grid)

    // look right
    x_hit := 0
    for i := x; i < max_x; i++ {
        if (*grid)[y][i] == 'X' {
            x_hit++
            break
        }
    }
    // look left
    for i := x; i >= min_x; i-- {
        if (*grid)[y][i] == 'X' {
            x_hit++
            break
        }
    }
    // look up
    for i := y; i >= min_y; i-- {
        if (*grid)[i][x] == 'X' {
            x_hit++
            break
        }
    }
    // look down
    for i := y; i < max_y; i++ {
        if (*grid)[i][x] == 'X' {
            x_hit++
            break
        }
    }
    if x_hit == 4 {
        (*grid)[y][x] = 'I'
        return
    }
    (*grid)[y][x] = 'O'
}

func isInGrid(x, y int, grid *[][]rune) bool {
    max_x := len((*grid)[0])

    x_hit := 0
    for i := x; i < max_x; i++ {
        if (*grid)[y][i] == 'W' {
            x_hit++
            // if i + 1 == max_x {
            //     x_hit++
            // }
        }
    }
    return x_hit % 2 == 1
}

func Day10Part2(filename string) (int, error) {
    lines, err := utils.ReadFileLines(filename)
    if err != nil {
        return -1, err
    }

    grid := getCharGrid(lines)
    g := buildGraph(grid)
    walkGraph(g, g.GetStartingVertex(), &grid)
    for y, row := range grid {
        for x, v := range row {
            if v == 'X' || v == 'W' {
                continue
            }
            if isInGrid(x, y, &grid) {
                grid[y][x] = 'I'
                continue
            }
            grid[y][x] = 'O'
        }
    }

    sum := 0
    for _, row := range grid {
        for _, v := range row {
            if v == 'I' {
                sum++
            }
        }
    }

    return sum, nil
}
