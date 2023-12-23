from enum import Enum
from string import ascii_lowercase

class BrickType(Enum):
    H = 0
    V = 1


class Brick:
    def __init__(self, id: int, x1: int, y1: int, z1: int, x2: int, y2: int, z2: int):
        self.id = id
        self.height = abs(z2 - z1) + 1
        self.brick_type = BrickType.H if self.height == 1 else BrickType.V
        self.x = range(x1, x2 + 1)
        self.y = range(y1, y2 + 1)
        self.x1 = x1
        self.x2 = x2
        self.y1 = y1
        self.y2 = y2
        self.z1 = z1
        self.z2 = z2
        self.supports = set()
        self.supported_by = set()


    def bricks_overlap(self, other):
        return max(self.x1, other.x1) <= min(self.x2, other.x2) and max(self.y1, other.y1) <= min(self.y2, other.y2)

    def __repr__(self):
        t = 'H' if self.brick_type == BrickType.H else 'V'
        return f'<Brick {self.id}>'


def get_stack(bricks):
    stack = [[bricks[0]]]
    for b in bricks[1:]:
        overlap = False
        for i, s in reversed(list(enumerate(stack))):
            for ss in s:
                if ss.bricks_overlap(b):
                    overlap = True
                    if ss != b:
                        ss.supports.add(b)
                        b.supported_by.add(ss)

            if overlap == True:
                if i + 1 == len(stack):
                    stack.append([b])
                else:
                    stack[i + 1].append(b)
                break
        if overlap == False:
            stack[0].append(b)

    return stack

def get_removable_bricks(stack):
    can_be_removed = set()

    for s in stack:
        for ss in s:
            counter = 0
            for b in ss.supports:
                if len(b.supported_by) > 1:
                    counter += 1
            if counter == len(ss.supports):
                can_be_removed.add(ss.id)

    return can_be_removed

def part_one(filename):
    bricks = parse_input(filename)

    stack = get_stack(bricks)

    can_be_removed = get_removable_bricks(stack)

    return len(can_be_removed)


def part_two(filename):
    bricks = parse_input(filename)
    stack = get_stack(bricks)


def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip().split('\n')

    bricks = []
    id = 0
    for i, l in enumerate(lines):
        brick_str = l.split('~')
        x1, y1, z1 = brick_str[0].split(',')
        x2, y2, z2 = brick_str[1].split(',')

        new_brick = Brick(id, int(x1), int(y1), int(z1), int(x2), int(y2), int(z2))
        for i in range(new_brick.height):
            bricks.append(new_brick)

        id += 1

    bricks.sort(key=lambda x: x.z1)

    return bricks


def main():
    example = './inputs/day_22_example.txt'
    test = './inputs/day_22_input.txt'

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    # print(f'Part Two Example: {part_two(example)}')
    # print(f'Part Two Input: {part_two(test)}')


if __name__ == '__main__':
    main()
