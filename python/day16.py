def travel(grid, start_x, start_y, start_dir):
    energized_grid = [[0 for _ in range(len(grid[0]))] for _ in range(len(grid))]

    queue = []
    queue.append((start_x, start_y, start_dir))
    cache = {}
    while len(queue) != 0:
        x, y, direction = queue.pop(0)
        if x < 0 or x >= len(grid[0]) or y < 0 or y >= len(grid):
            continue

        key = (x, y, direction)
        if key in cache:
            continue

        cache[key] = True

        energized_grid[y][x] += 1    
        char = grid[y][x]

        match char:
            case '.':
                if direction == 'R':
                    queue.append((x + 1, y, 'R'))
                elif direction == 'L':
                    queue.append((x - 1, y, 'L'))
                elif direction == 'U':
                    queue.append((x, y - 1, 'U'))
                elif direction == 'D':
                    queue.append((x, y + 1, 'D'))
            case '|':
                if direction == 'R' or direction == 'L':
                    queue.append((x, y - 1, 'U'))
                    queue.append((x, y + 1, 'D'))
                elif direction == 'U':
                    queue.append((x, y - 1, 'U'))
                elif direction == 'D':
                    queue.append((x, y + 1, 'D'))
            case '-':
                if direction == 'U' or direction == 'D':
                    queue.append((x + 1, y, 'R'))
                    queue.append((x - 1, y, 'L'))
                elif direction == 'R':
                    queue.append((x + 1, y, 'R'))
                elif direction == 'L':
                    queue.append((x - 1, y, 'L'))
            case '\\':
                if direction == 'R':
                    queue.append((x, y + 1, 'D'))
                elif direction == 'L':
                    queue.append((x, y - 1, 'U'))
                elif direction == 'U':
                    queue.append((x - 1, y, 'L'))
                elif direction == 'D':
                    queue.append((x + 1, y, 'R'))
            case '/':
                if direction == 'R':
                    queue.append((x, y - 1, 'U'))
                elif direction == 'L':
                    queue.append((x, y + 1, 'D'))
                elif direction == 'U':
                    queue.append((x + 1, y, 'R'))
                elif direction == 'D':
                    queue.append((x - 1, y, 'L'))


    return energized_grid

def part_one(filename):
    grid = parse_input(filename)

    energized_grid = travel(grid, 0, 0, 'R')

    zeros = 0
    for i in energized_grid:
        zeros += i.count(0)

    # for i in energized_grid:
    #     print(''.join(['.' if j == 0 else '#' for j in i]))

    acc = len(grid) * len(grid) - zeros

    return acc


def score(grid):
    acc = 0
    for i in grid:
        acc += i.count(0)
    return len(grid) ** 2 - acc


def part_two(filename):
    grid = parse_input(filename)
    
    accs = []
    # accross top
    for i in range(len(grid[0])):
        accs.append(score(travel(grid, i, 0, 'D')))
        accs.append(score(travel(grid, i, len(grid) - 1, 'U')))

    for i in range(len(grid)):
        accs.append(score(travel(grid, 0, i, 'R')))
        accs.append(score(travel(grid, len(grid[0]) - 1, i, 'L')))

    return max(accs)

def parse_input(filename):
    with open(filename, 'r') as f:
        return [line.strip() for line in f.readlines()]


def main():
    example = './inputs/day_16_example.txt'
    test = './inputs/day_16_input.txt'

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example: {part_two(example)}')
    print(f'Part Two Input: {part_two(test)}')


if __name__ == "__main__":
    main()
