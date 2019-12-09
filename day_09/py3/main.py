from collections import defaultdict
from itertools import permutations
import re


def get_input() -> str:
    with open("day_09/input.txt") as f:
        return f.read()


def get_input_lines() -> list:
    return get_input().split("\n")


def ints(text: str) -> tuple:
    return tuple(map(int, re.findall("([+\-0-9]+)", text)))


def flat_map(collection):
    for container in collection:
        for element in container:
            yield element


class IntCode:
    """A basic class to execute intcode program."""

    def __init__(self, values: list, inputs: list):
        self.code = defaultdict(int, [(i, values[i]) for i in range(len(values))])
        self.pointer = 0
        self.inputs = inputs
        self.outputs = []
        self.running = True
        self.paused = False
        self.rel_base = 0
        self.modes = [0, 0, 0]

    def tick(self):
        # Calculate all instruction flags
        inst = self.code[self.pointer]
        opcode = inst % 100
        self.modes = [(inst // 100) % 10, (inst // 1000) % 10, (inst // 10000) % 10]

        self.paused = False
        if opcode == 1:
            self.code[self.addr(3)] = self.arg(1) + self.arg(2)
            self.pointer += 4
        elif opcode == 2:
            self.code[self.addr(3)] = self.arg(1) * self.arg(2)
            self.pointer += 4
        elif opcode == 3:
            if len(self.inputs) > 0:
                self.code[self.addr(1)] = self.inputs.pop(0)
                self.pointer += 2
            else:
                self.paused = True
        elif opcode == 4:
            self.outputs.append(self.arg(1))
            self.pointer += 2
        elif opcode == 5:
            if self.arg(1) != 0:
                self.pointer = self.arg(2)
            else:
                self.pointer += 3
        elif opcode == 6:
            if self.arg(1) == 0:
                self.pointer = self.arg(2)
            else:
                self.pointer += 3
        elif opcode == 7:
            self.code[self.addr(3)] = 1 if self.arg(1) < self.arg(2) else 0
            self.pointer += 4
        elif opcode == 8:
            self.code[self.addr(3)] = 1 if self.arg(1) == self.arg(2) else 0
            self.pointer += 4
        elif opcode == 9:
            self.rel_base += self.arg(1)
            self.pointer += 2
        elif opcode == 99:
            self.running = False

    def run(self):
        while self.running:
            self.tick()
        return self

    def arg(self, i: int) -> int:
        """ Internal function to get an argument, after opcode / position flags have been calculated
        :param i: the index of the argument to get, as an offset from the opcode
        :return: the argument value (immediate, positional, or relative)
        """
        if self.modes[i - 1] == 0:
            return self.code[self.code[self.pointer + i]]
        elif self.modes[i - 1] == 1:
            return self.code[self.pointer + i]
        elif self.modes[i - 1] == 2:
            return self.code[self.rel_base + self.code[self.pointer + i]]
        raise ValueError("Unknown argument mode %d" % self.modes[i - 1])

    def addr(self, i: int) -> int:
        """ Internal function to get an address, after opcode / position flags have been calculated
        :param i: the index of the address to get, as an offset from the opcode
        :return: the address (positional, or relative, addresses can't be immediate)
        """
        if self.modes[i - 1] == 0:
            return self.code[self.pointer + i]
        elif self.modes[i - 1] == 1:
            raise ValueError("Tried to get address with immediate argument mode")
        elif self.modes[i - 1] == 2:
            return self.rel_base + self.code[self.pointer + i]
        raise ValueError("Unknown argument mode %d" % self.modes[i - 1])


def part1(code: list):
    return IntCode(input_code, [1]).run().outputs[0]


def part2(code: list):
    return IntCode(code, [2]).run().outputs[0]


if __name__ == "__main__":
    input_code = [*ints(get_input())]
    print("Part 1:", part1(input_code))
    print("Part 2:", part2(input_code))
