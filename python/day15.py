import re

def hash_seq(s):
    cv = 0
    for c in s:
        ascii_val = ord(c)
        cv += ascii_val
        cv *= 17
        cv %= 256

    return cv


def part_one(filename):
    seqs = parse_input(filename)

    acc = 0
    for s in seqs:
        acc += hash_seq(s)

    return acc 


def part_two(filename):
    seqs = parse_input(filename)
    seqs = [ re.sub('-', ' - ', s) for s in seqs ]
    seqs = [ re.sub('=', ' = ', s) for s in seqs ]
    seqs = [ tuple(re.split(' ', s)) for s in seqs ]

    boxes = [ [] for _ in range(256) ]

    for s in seqs:
        hash_val = hash_seq(s[0])
        label = s[0]
        op = s[1]
        lense = s[2]

        in_box = False
        for l in boxes[hash_val]:
            if l[0] == s[0]:
                in_box = True
                break

        if op == '-' and in_box:
            for i in range(len(boxes[hash_val])):
                if boxes[hash_val][i][0] == label:
                    boxes[hash_val].pop(i)
                    break
        elif op == '=': 
            if not in_box:
                boxes[hash_val].append((label, lense))
                continue
            for i in range(len(boxes[hash_val])):
                if boxes[hash_val][i][0] == label:
                    boxes[hash_val][i] = (label, lense)
                    break

    acc = 0
    for i, b in enumerate(boxes):
        for j, l in enumerate(b):
            acc += (i + 1) * (j + 1) * int(l[1])

    return acc


def parse_input(filename):
    with open(filename, 'r') as f:
        line = f.read().strip()

    return line.split(',')


def main():
    example = './inputs/day_15_example.txt'
    test = './inputs/day_15_input.txt'

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example: {part_two(example)}')
    print(f'Part Two Input: {part_two(test)}')


if __name__ == "__main__":
    main()
