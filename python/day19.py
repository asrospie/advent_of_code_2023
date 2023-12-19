import re


def part_access(part, value):
    if value == 'x':
        return part[0]
    elif value == 'm':
        return part[1]
    elif value == 'a':
        return part[2]
    elif value == 's':
        return part[3]
    return 0


def parse_rule(workflow, part):
    rules = workflow.split(',')
    last_state = None
    for rule in rules:
        if '<' in rule:
            cat, test = rule.split('<')
            test_value, next_state = test.split(':')
            part_cat = part_access(part, cat)
            if part_cat < int(test_value):
                return next_state
        elif '>' in rule:
            cat, test = rule.split('>')
            test_value, next_state = test.split(':')
            part_cat = part_access(part, cat)
            if part_cat > int(test_value):
                return next_state
        last_state = rule
    return last_state


def get_workflows(workflows_str):
    workflows = {}
    for w in workflows_str:
        regex = r'(\w+){(.*)}'
        name, rules = re.findall(regex, w)[0]
        workflows[name] = rules
    return workflows

def get_parts(parts_str):
    parts = []
    for p in parts_str:
        regex = r'\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}'
        x, m, a, s = re.findall(regex, p)[0]
        parts.append((int(x), int(m), int(a), int(s)))
    return parts


def part_one(filename):
    workflows_str, parts_str = parse_input(filename)
    workflows = get_workflows(workflows_str)
    parts = get_parts(parts_str)

    accepted = 0
    acc = 0
    for part in parts:
        state = 'in'
        while state != 'A' or state != 'R':
            if state == 'A' or state == 'R':
                break
            workflow = workflows[state]
            state = parse_rule(workflow, part)
        if state == 'A':
            accepted += 1
            acc += sum(part)

    return acc


def part_two(filename):
    workflows_str, _ = parse_input(filename)
    workflows = get_workflows(workflows_str)

    return 0


def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip()

    workflows, parts = lines.split('\n\n')

    workflows = workflows.split('\n')
    parts = parts.split('\n')

    return workflows, parts


def main():
    example = './inputs/day_19_example.txt'
    test = './inputs/day_19_input.txt'

    print(f'Part One Example: {part_one(example)}')
    print(f'Part One Input: {part_one(test)}')

    print(f'Part Two Example: {part_two(example)}')
    # print(f'Part Two Input: {part_two(test)}')


if __name__ == '__main__':
    main()
