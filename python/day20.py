import re
import json
from enum import Enum

# Create Enum
class Pulse(Enum):
    HIGH = 1
    LOW = 0


class Broadcaster:
    def __init__(self, outputs):
        self.outputs = outputs


class FlipFlop:
    def __init__(self, tag, outputs):
        self.tag = tag
        self.state = False
        self.outputs = outputs

    def receive_pulse(self, pulse):
        if pulse == Pulse.LOW:
            if self.state == True:
                self.state = False
                return Pulse.LOW
            elif self.state == False:
                self.state = True
                return Pulse.HIGH
        return None


class Conjunction:
    def __init__(self, tag, outputs):
        self.tag = tag
        self.outputs = outputs
        self.inputs = {}


    def add_input(self, key):
        self.inputs[key] = Pulse.LOW

    
    def receive_pulse(self, pulse, from_tag):
        self.inputs[from_tag] = pulse

        if Pulse.LOW in self.inputs.values():
            return Pulse.HIGH

        return Pulse.LOW


def part_one(filename, btn_presses=1000):
    lines = parse_input(filename)

    config = gather_config(lines)
    low_pulses = 0
    high_pulses = 0
    

    for i in range(btn_presses):
        queue = [(Pulse.LOW, 'broadcaster', 'btn')]
        while len(queue) > 0:
            pulse, tag, prev = queue.pop(0)
            if pulse == Pulse.LOW:
                low_pulses += 1
            elif pulse == Pulse.HIGH:
                high_pulses += 1

            if tag not in config:
                continue

            send_pulse = pulse
            module = config[tag]
            outputs = module.outputs

            if type(module) == Conjunction:
                send_pulse = module.receive_pulse(pulse, prev) 
            elif type(module) == FlipFlop:
                send_pulse = module.receive_pulse(pulse)

            if send_pulse == None:
                continue
            
            for output in outputs:
                queue.append((send_pulse, output, tag))

    return low_pulses * high_pulses


def part_two(filename):
    lines = parse_input(filename)

    config = gather_config(lines)
    low_pulses = 0
    high_pulses = 0
    

    btn_presses = 0
    rx_found = False
    while not rx_found:
        queue = [(Pulse.LOW, 'broadcaster', 'btn')]
        btn_presses += 1
        print(f'btn presses: {btn_presses}')
        while len(queue) > 0:
            pulse, tag, prev = queue.pop(0)
            if pulse == Pulse.LOW:
                low_pulses += 1
            elif pulse == Pulse.HIGH:
                high_pulses += 1

            if tag == 'rx' and pulse == Pulse.LOW:
                rx_found = True
                break

            if tag not in config:
                continue

            send_pulse = pulse
            module = config[tag]
            outputs = module.outputs

            if type(module) == Conjunction:
                send_pulse = module.receive_pulse(pulse, prev) 
            elif type(module) == FlipFlop:
                send_pulse = module.receive_pulse(pulse)

            if send_pulse == None:
                continue
            
            for output in outputs:
                queue.append((send_pulse, output, tag))

    return low_pulses * high_pulses


def gather_config(lines):
    config = {}
    for l in lines:
        regex = r'([%&]*)(\w+) -> ([\w\s,]+)'
        matches = re.findall(regex, l)
        module, tag, output = matches[0]

        if module == '%':
            config[tag] = FlipFlop(tag, output.split(', '))
        elif module == '&':
            config[tag] = Conjunction(tag, output.split(', '))
        elif module == '':
            config[tag] = Broadcaster(output.split(', '))

    for key, value in config.items():
        module = value
        if type(module) == Conjunction:
            for k, v in config.items():
                if module.tag in v.outputs:
                    module.add_input(k)


    return config


def parse_input(filename):
    with open(filename) as f:
        lines = f.read().strip().split('\n')

    return lines


def main():
    example_one = './inputs/day_20_example_1.txt'
    example_two = './inputs/day_20_example_2.txt'
    test = './inputs/day_20_input.txt'

    # print(f'Part One Example 1: {part_one(example_one)}')
    # print(f'Part One Example 2: {part_one(example_two)}')
    # print(f'Part One Input: {part_one(test)}')

    # print(f'Part Two Example 1: {part_two(example_one)}')
    # print(f'Part Two Example 2: {part_two(example_two)}')
    print(f'Part Two Input: {part_two(test)}')


if __name__ == '__main__':
    main()
