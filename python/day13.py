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
        return reflection_idx

    return 0


def is_off_by_one(str1, str2):
    if len(str1) != len(str2):
        return False

    off = 0
    for i in range(len(str1)):
        if str1[i] != str2[i]:
            off += 1

    return off == 1


def find_smudge(pattern):
    stack = []

    reflection_idx = 0
    reflection = False
    reflections = []
    for i, line in enumerate(pattern):
        stack.append(line)

        if i + 1 >= len(pattern):
            break

        same = pattern[i + 1] == line
        off_by_one = is_off_by_one(line, pattern[i + 1])
        off_by_one_count = 1 if off_by_one else 0
        
        stack_idx = len(stack) - 2
        if same or off_by_one:
            reflection = True
            # print(f'Found mirror at {i + 1}')
            reflection_idx = i + 1 

            j = i + 2
            while stack_idx >= 0 and j < len(pattern):
                check = stack[stack_idx]
                off = is_off_by_one(check, pattern[j])
                if pattern[j] != check and not off:
                    reflection = False
                    break
                if off:
                    off_by_one_count += 1
                j += 1
                reflection = True
                stack_idx -= 1

            if reflection:
                reflections.append((reflection_idx, off_by_one_count))

    if len(reflections) > 0:
        for r in reflections:
            if r[1] == 1:
                return r[0]
    return 0


def part_one(filename):
    content = parse_input(filename)

    sum = 0
    for pattern in content:
        sum += find_matches(rotate_list(pattern))
        sum += find_matches(pattern) * 100

    return sum


def part_two(filename):
    content = parse_input(filename)

    sum = 0
    for pattern in content:
        sum += find_smudge(rotate_list(pattern))
        sum += find_smudge(pattern) * 100
    return sum


def main():
    example = './inputs/day_13_example.txt'
    test = './inputs/day_13_input.txt'
    # print(f'Part One:\t{part_one(example)}')
    # print(f'Part One:\t{part_one(test)}')

    print(f'\nPart Two:\t{part_two(example)}')
    print(f'Part Two:\t{part_two(test)}')

if __name__ == "__main__":
    main()
