def simulate(lines):
    last_changes = 1
    while last_changes != 0:
        last_changes = 0
        for i in range(1, len(lines)):
            for j, c in enumerate(lines[i]):
                if c == 'O' and lines[i-1][j] == '.':
                    lines[i-1][j] = 'O'
                    lines[i][j] = '.'
                    last_changes += 1

    return lines


def rotate_right(lines):
    new_lines = []
    for i in range(len(lines)):
        new_lines.append([])
        for j in range(len(lines[0])):
            new_lines[i].append(lines[len(lines) - j - 1][i])

    return new_lines


def cycle(lines, cycles):
    cycle_found = False
    cache = {}
    i = 0
    while i < cycles:
        # north
        lines = simulate(lines)
        # west
        lines = rotate_right(lines)
        lines = simulate(lines)
        # south
        lines = rotate_right(lines)
        lines = simulate(lines)
        # east
        lines = rotate_right(lines)
        lines = simulate(lines)
        # north
        lines = rotate_right(lines)

        key = ''.join([ ''.join(x) for x in lines ])
        if not cycle_found and key in cache:
            cycle_found = True
            i = cycles - ((cycles - i) % (i - cache[key]))

        if not cycle_found:
            cache[key] = i

        i += 1

    return lines

def find_sum(grid):
    sum = 0
    for i, l in enumerate(grid):
        sum += (len(grid) - i) * l.count('O')
    return sum


def parse_input(filename):
    with open(filename) as f:
        lines = [ x.strip() for x in f.readlines() ]

    return [ list(x) for x in lines ]


def part_one(filename):
    grid = parse_input(filename)

    return find_sum(simulate(grid))


def part_two(filename):
    grid = parse_input(filename)

    return find_sum(cycle(grid, 1_000_000_000))


def main():
    example = "inputs/day_14_example.txt"
    test = "inputs/day_14_input.txt"

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example: {part_two(example)}')
    print(f'Part Two Input: {part_two(test)}')

if __name__ == "__main__":
    main()
