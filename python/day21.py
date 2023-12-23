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


def walk(grid, pos, max_steps=6):
    walked = set()
    queue = [(pos, 0)]

    possibilities = set()
    walked.add(pos)

    while len(queue) > 0:
        pos, steps = queue.pop(0)
        steps_remaining = max_steps - steps
        x, y = pos

        # for w in walked:
        #     if w not in possibilities and manhattan_distance(w, pos) == steps_remaining:
        #         wx, wy = w
        #         grid[wy][wx] = 'O'
        #         possibilities.add(w)
        if steps % 2 == 0:
            possibilities.add(pos)

        if steps == max_steps - 1:
            possibilities.add(pos)
            continue

        for n in get_neighbors(grid, pos):
            if n not in walked:
                walked.add(n)
                queue.append((n, steps + 1))

    return len(possibilities)


def part_one(filename, max_steps=6):
    grid = parse_input(filename)

    starting_pos = (0, 0)
    for y, row in enumerate(grid):
        for x, v in enumerate(row):
            if v == 'S':
                starting_pos = (x, y)
                break

    return walk(grid, starting_pos, max_steps)


def part_two(filename):
    grid = parse_input(filename)

    start = (0, 0)
    for y, row in enumerate(grid):
        for x, v in enumerate(row):
            if v == 'S':
                start = (x, y)
                break

    size = len(grid)
    max_steps = 26501365
    grid_width = (max_steps // size) - 1
    odd_grids = (grid_width // 2 * 2 + 1) ** 2
    even_grids = ((grid_width + 1) // 2 * 2) ** 2

    odd_points = part_one(filename, size * 2 + 1)
    even_points = part_one(filename, size * 2)

    top_corner = walk(grid, (start[0], size - 1), size - 1)
    right_corner = walk(grid, (0, start[1]), size - 1)
    bottom_corner = walk(grid, (start[0], 0), size - 1)
    left_corner = walk(grid, (size - 1, start[1]), size - 1)

    small_tr = walk(grid, (0, size - 1), size // 2 - 1)
    small_tl = walk(grid, (size - 1, size - 1), size // 2 - 1)
    small_br = walk(grid, (0, 0), size // 2 - 1)
    small_bl = walk(grid, (size - 1, 0), size // 2 - 1)

    large_tr = walk(grid, (0, size - 1), size * 3 // 2 - 1)
    large_tl = walk(grid, (size - 1, size - 1), size * 3 // 2 - 1)
    large_br = walk(grid, (0, 0), size * 3 // 2 - 1)
    large_bl = walk(grid, (size - 1, 0), size * 3 // 2 - 1)

    acc = \
            odd_grids * odd_points + \
            even_grids * even_points + \
            top_corner + bottom_corner + left_corner + right_corner + \
            (grid_width + 1) * (small_tr + small_tl + small_br + small_bl) + \
            grid_width * (large_tr + large_tl + large_br + large_bl)

    return acc

def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip().split('\n')

    return [ [ c for c in line ] for line in lines ]


def main():
    example = './inputs/day_21_example.txt'
    test = './inputs/day_21_input.txt'

    print(f'Part One Example: {part_one(example)}')
    # print(f'Part One Input: {part_one(test, 64)}')

    print(f'Part Two Input: {part_two(test)}')


if __name__ == '__main__':
    main()
