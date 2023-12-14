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

    sum = 0
    for i, l in enumerate(lines):
        sum += (len(lines) - i) * l.count('O') 

    return sum


def simulate_cycle(lines):
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
        for j in range(len(lines)):
            new_lines[i].append(lines[j][i])

    return new_lines


def cycle(lines, cycles):
    for i in range(cycles):
        print(f'Cycle #{i}')
        # north
        lines = simulate_cycle(lines)
        # west
        lines = rotate_right(lines)
        lines = simulate_cycle(lines)
        # south
        lines = rotate_right(lines)
        lines = simulate_cycle(lines)
        # east
        lines = rotate_right(lines)
        lines = simulate_cycle(lines)
        # north
        lines = rotate_right(lines)

    sum = 0
    for i, l in enumerate(lines):
        sum += (len(lines) - i) * l.count('O')
    return sum


def part_one(filename):
    with open(filename) as f:
        lines = [ x.strip() for x in f.readlines() ]
        lines = [ list(x) for x in lines ]

    return simulate(lines)


def part_two(filename):
    with open(filename) as f:
        lines = [ x.strip() for x in f.readlines() ]
        lines = [ list(x) for x in lines ]

    return cycle(lines, 1_000_000_000)


def main():
    example = "inputs/day_14_example.txt"
    test = "inputs/day_14_input.txt"

    # print(f'Part One Example: {part_one(example)}')
    # print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example: {part_two(example)}')

if __name__ == "__main__":
    main()
