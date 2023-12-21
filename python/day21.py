def get_neighbors(grid, pos):
    x, y = pos
    neighbors = []

    if x > 0 and grid[y][x - 1] != '#':
        neighbors.append((x - 1, y))
    if x < len(grid[0]) - 1 and grid[y][x + 1] != '#':
        neighbors.append((x + 1, y))
    if y > 0 and grid[y - 1][x] != '#':
        neighbors.append((x, y - 1))
    if y < len(grid) - 1 and grid[y + 1][x] != '#':
        neighbors.append((x, y + 1))

    return neighbors


def manhattan_distance(a, b):
    ax, ay = a
    bx, by = b

    return abs(ax - bx) + abs(ay - by)


def part_one(filename, max_steps=6):
    grid = parse_input(filename)

    starting_pos = (0, 0)

    for y, row in enumerate(grid):
        for x, v in enumerate(row):
            if v == 'S':
                starting_pos = (x, y)
                break

    walked = set()
    queue = [(starting_pos, 0)]

    possibilities = set()
    walked.add(starting_pos)

    while len(queue) > 0:
        pos, steps = queue.pop(0)
        steps_remaining = max_steps - steps
        x, y = pos

        for w in walked:
            if w not in possibilities and manhattan_distance(w, pos) == steps_remaining:
                wx, wy = w
                grid[wy][wx] = 'O'
                possibilities.add(w)

        if steps == max_steps:
            grid[y][x] = 'O'
            possibilities.add(pos)
            continue

        for n in get_neighbors(grid, pos):
            if n not in walked:
                walked.add(n)
                queue.append((n, steps + 1))

    return len(possibilities)


def part_two(filename):
    pass


def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip().split('\n')

    return [ [ c for c in line ] for line in lines ]


def main():
    example = './inputs/day_21_example.txt'
    test = './inputs/day_21_input.txt'

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test, 64)}')

    # print(f'Part Two Example: {part_two(example)}')
    # print(f'Part Two Input: {part_two(test)}')


if __name__ == '__main__':
    main()
