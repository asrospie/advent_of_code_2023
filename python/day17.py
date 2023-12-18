from heapq import heappush, heappop


def dijkstra_modified(grid, start, end):
    seen = set()
    pq = []
    # (heat loss, x, y, dx, dy, consecutive)
    heappush(pq, (0, start[0], start[1], 0, 0, 0))

    while pq:
        loss, x, y, dx, dy, consecutive = heappop(pq)

        if (x, y) == end:
            return loss

        if (x, y, dx, dy, consecutive) in seen:
            continue

        seen.add((x, y, dx, dy, consecutive))

        if consecutive < 3 and (dx, dy) != (0, 0):
            x_next = x + dx
            y_next = y + dy
            if 0 <= x_next < len(grid[0]) and 0 <= y_next < len(grid):
                heappush(pq, (loss + grid[y_next][x_next], x_next, y_next, dx, dy, consecutive + 1))

        for x_dir_next, y_dir_next in [ (0, 1), (0, -1), (1, 0), (-1, 0) ]:
            if (x_dir_next, y_dir_next) == (dx, dy) or (x_dir_next, y_dir_next) == (-dx, -dy):
                continue

            x_next = x + x_dir_next
            y_next = y + y_dir_next

            if 0 <= x_next < len(grid[0]) and 0 <= y_next < len(grid):
                heappush(pq, (loss + grid[y_next][x_next], x_next, y_next, x_dir_next, y_dir_next, 1))

    return 0


def dijkstra_modified_ten(grid, start, end):
    seen = set()
    pq = []
    # (heat loss, x, y, dx, dy, consecutive)
    heappush(pq, (0, start[0], start[1], 0, 0, 0))

    while pq:
        loss, x, y, dx, dy, consecutive = heappop(pq)

        if (x, y) == end:
            return loss

        if (x, y, dx, dy, consecutive) in seen:
            continue

        seen.add((x, y, dx, dy, consecutive))

        if consecutive < 10 and (dx, dy) != (0, 0):
            x_next = x + dx
            y_next = y + dy
            if 0 <= x_next < len(grid[0]) and 0 <= y_next < len(grid):
                heappush(pq, (loss + grid[y_next][x_next], x_next, y_next, dx, dy, consecutive + 1))

        for x_dir_next, y_dir_next in [ (0, 1), (0, -1), (1, 0), (-1, 0) ]:
            if (x_dir_next, y_dir_next) == (dx, dy) or (x_dir_next, y_dir_next) == (-dx, -dy):
                continue

            x_next = x + x_dir_next
            y_next = y + y_dir_next

            if 0 <= x_next < len(grid[0]) and 0 <= y_next < len(grid) and (consecutive >= 4 or len(seen) < 4):
                heappush(pq, (loss + grid[y_next][x_next], x_next, y_next, x_dir_next, y_dir_next, 1))

    return 0


def part_one(filename):
    grid = parse_input(filename)

    return dijkstra_modified(grid, (0, 0), (len(grid[0]) - 1, len(grid) - 1))


def part_two(filename):
    grid = parse_input(filename)

    return dijkstra_modified_ten(grid, (0, 0), (len(grid[0]) - 1, len(grid) - 1))


def parse_input(filename):
    with open(filename) as f:
        return [ [ int(c) for c in line.strip() ] for line in f.readlines() ]


def main():
    example = './inputs/day_17_example.txt'
    test = './inputs/day_17_input.txt'

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example:\t{part_two(example)}')
    print(f'Part Two Input:\t{part_two(test)}')

if __name__ == '__main__':
    main()
