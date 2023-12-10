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
    fmt.Println("Adding edge")
    // check if can connect
    if !(v2.CanConnectDestination(search_dir) && v1.CanConnectSource(search_dir)) {
        fmt.Printf("Cannot connect %s to %s from %s\n", typeToString(v1.Type), typeToString(v2.Type), dirToString(search_dir))
        return
    }
    fmt.Printf("Connecting %s to %s from %s\n", typeToString(v1.Type), typeToString(v2.Type), dirToString(search_dir))
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
        fmt.Println("==================================")
        fmt.Printf("Current: %v, Prev: %v\n", current.String(), prev.String())
        next := current.Edges[0].Vertex
        fmt.Printf("Next: %v\n", next.String())
        next_is_prev := next.IsEqual(prev)
        fmt.Println(next_is_prev)
        if next_is_prev {
            next = current.Edges[1].Vertex
            fmt.Printf("Next: %v\n", next.String())
        }
        fmt.Println(next.String())
        prev = current
        current = next
        counter++
        fmt.Println("==================================")
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
    fmt.Println(g.String())
    v := g.GetStartingVertex()
    return g.FindFurthestVertexDistance(v), nil
    // return -1, nil
}

func Day10Part2(filename string) (int, error) {
    return -1, nil
}
