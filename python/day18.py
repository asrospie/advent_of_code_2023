import re

# brute forced this one, but it works
def part_one_brute(filename):
    dig_plan = parse_input(filename)

    grid = [['.' for _ in range(1000)] for _ in range(1000)]

    p = (500, 500)


    for dir, m, _ in dig_plan:
        x = p[0]
        y = p[1]
        if dir == 'U':
            for i in range(0, m+1):
                grid[y-i][x] = 'U'
                p = (x, y-i)
        if dir == 'D':
            for i in range(0, m+1):
                grid[y+i][x] = 'D'
                p = (x, y+i)
        if dir == 'L':
            for i in range(1, m+1):
                grid[y][x-i] = '#'
                p = (x-i, y)
        if dir == 'R':
            for i in range(1, m+1):
                grid[y][x+i] = '#'
                p = (x+i, y)

    for y, row in enumerate(grid):
        filling = False
        for x, c in enumerate(row):
            if c == 'U' and not filling:
                filling = True
            elif c == 'D' and filling:
                filling = False
            elif c == '.' and filling:
                grid[y][x] = '#'

    acc = 0
    for row in grid:
        acc += row.count('#')
        acc += row.count('U')
        acc += row.count('D')

    return acc


def part_one(filename):
    dig_plan = parse_input_basic(filename)

    return shoelace_pick(dig_plan)


def determinant(p1, p2):
    x1, y1 = p1
    x2, y2 = p2

    return (x1 * y2) - (x2 * y1)


def shoelace_pick(dig_plan):
    cur_pos = (0, 0)
    points = [cur_pos]
    permimeter = 0
    for row in dig_plan:
        d, m = row
        p = cur_pos
        if d == 'R':
            p = (p[0] + m, p[1])
        elif d == 'L':
            p = (p[0] - m, p[1])
        elif d == 'U':
            p = (p[0], p[1] - m)
        elif d == 'D':
            p = (p[0], p[1] + m)
        cur_pos = p
        points.append(cur_pos)
        permimeter += m

    # check if need to reverse
    first_lateral_move = None
    for i in range(len(dig_plan)):
        if dig_plan[i][0] == 'R' or dig_plan[i][0] == 'L':
            first_lateral_move = dig_plan[i][0]
            break
    if first_lateral_move == 'R':
        points = points[::-1]

    acc = 0
    for i in range(0, len(points), 2):
        if i + 1 >= len(points):
            break
        p1 = points[i]
        p2 = points[i+1]
        acc += determinant(p1, p2)
    acc += determinant(points[-1], points[0])

    return abs(acc) + permimeter // 2 + 1


# go figure, brute force doesn't work for part two
def part_two(filename):
    dig_plan = parse_input_w_hex(filename)

    return shoelace_pick(dig_plan)


def parse_input_w_hex(filename):
    dig_plan = parse_input(filename)
    for i in range(len(dig_plan)):
        _, _, m = dig_plan[i]
        hex_string = re.findall(r'[\dA-Za-z]+', m[:-1])[0]
        dir = hex_string[-1]
        dir = {
            '0': 'R',
            '1': 'D',
            '2': 'L',
            '3': 'U',
        }.get(dir, None)
        dig_plan[i] = (dir, int(hex_string[:-1], 16))

    return dig_plan


def parse_input_basic(filename):
    dig_plan = parse_input(filename)
    for i in range(len(dig_plan)):
        dig_plan[i] = (dig_plan[i][0], int(dig_plan[i][1]))

    return dig_plan


def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip().split('\n')

    dig_plan = []
    for line in lines:
        dir, m, color = line.split(' ')
        dig_plan.append((dir, int(m), color))

    return dig_plan


def main():
    example = './inputs/day_18_example.txt'
    test = './inputs/day_18_input.txt'

    print(f'Part One Example (Brute Force): {part_one_brute(example)}')
    print(f'Part One Input (Brute Force): {part_one_brute(test)}')

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example: {part_two(example)} :: Expected: 952408144115')
    print(f'Part Two Input: {part_two(test)}')


if __name__ == '__main__':
    main()
