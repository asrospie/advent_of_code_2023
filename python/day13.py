def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip()

    return  [ line.split('\n') for line in lines.split('\n\n') ]


def rotate_list(pattern):
    num_rows = len(pattern)
    num_cols = len(pattern[0])

    rotated = [ '' for _ in range(num_cols) ]
    for i in range(num_rows):
        for j in range(num_cols):
            rotated[j] += pattern[i][j]

    return rotated


def find_matches(pattern) -> int:
    stack = []

    reflection_idx = 0
    reflection = False
    for i, line in enumerate(pattern):
        stack.append(line)

        # found mirror
        stack_idx = len(stack) - 2
        if i + 1 < len(pattern) and pattern[i + 1] == line:
            reflection = True
            # print(f'Found mirror at {i + 1}')
            reflection_idx = i + 1 

            j = i + 2
            while stack_idx >= 0 and j < len(pattern):
                # print(f'j: {j}, stack_idx: {stack_idx}')
                # print(f'Checking {pattern[j]} with {stack[stack_idx]}')
                check = stack[stack_idx]
                if pattern[j] != check:
                    reflection = False
                    break
                j += 1
                reflection = True
                stack_idx -= 1

            if reflection:
                break

    if reflection:
        # print('Found reflection')
        return reflection_idx

    return 0


def part_one(filename):
    content = parse_input(filename)

    sum = 0
    for pattern in content:
        sum += find_matches(rotate_list(pattern))
        sum += find_matches(pattern) * 100

    return sum


def part_two(filename):
    pass


def main():
    example = './inputs/day_13_example.txt'
    test = './inputs/day_13_input.txt'
    print("Part one:", part_one(example))
    print("Part one:", part_one(test))


if __name__ == "__main__":
    main()
