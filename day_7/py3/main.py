from itertools import permutations

inputs = [int(i) for i in open("day_7/input.txt").read().split(",")]
phase_setting_sequences = list(permutations(range(5), 5))


def runner(phase_setting):
    program = list(inputs)
    pos = 0
    next_input = 0
    while True:
        instructions = [int(x) for x in str(program[pos])]
        opcode = (
            0 if len(instructions) == 1 else instructions[-2]
        ) * 10 + instructions[-1]
        instructions = instructions[:-2]
        if opcode == 1:
            while len(instructions) < 3:
                instructions = [0] + instructions
            ins1, ins2, ins3 = program[pos + 1], program[pos + 2], program[pos + 3]
            program[ins3] = (ins1 if instructions[2] == 1 else program[ins1]) + (
                ins2 if instructions[1] == 1 else program[ins2]
            )
            pos += 4
        elif opcode == 2:
            while len(instructions) < 3:
                instructions = [0] + instructions
            ins1, ins2, ins3 = program[pos + 1], program[pos + 2], program[pos + 3]
            program[ins3] = (ins1 if instructions[2] == 1 else program[ins1]) * (
                ins2 if instructions[1] == 1 else program[ins2]
            )
            pos += 4
        elif opcode == 3:
            ins1 = program[pos + 1]
            program[ins1] = phase_setting[next_input]
            next_input += 1
            pos += 2
        elif opcode == 4:
            ins1 = program[pos + 1]
            return program[ins1]
            pos += 2
        elif opcode == 5:
            while len(instructions) < 2:
                instructions = [0] + instructions
            ins1, ins2 = program[pos + 1], program[pos + 2]
            if (ins1 if instructions[1] == 1 else program[ins1]) != 0:
                pos = ins2 if instructions[0] == 1 else program[ins2]
            else:
                pos += 3
        elif opcode == 6:
            while len(instructions) < 2:
                instructions = [0] + instructions
            ins1, ins2 = program[pos + 1], program[pos + 2]
            if (ins1 if instructions[1] == 1 else program[ins1]) == 0:
                pos = ins2 if instructions[0] == 1 else program[ins2]
            else:
                pos += 3
        elif opcode == 7:
            while len(instructions) < 3:
                instructions = [0] + instructions
            ins1, ins2, ins3 = program[pos + 1], program[pos + 2], program[pos + 3]
            if (ins1 if instructions[2] == 1 else program[ins1]) < (
                ins2 if instructions[1] == 1 else program[ins2]
            ):
                program[ins3] = 1
            else:
                program[ins3] = 0
            pos += 4
        elif opcode == 8:
            while len(instructions) < 3:
                instructions = [0] + instructions
            ins1, ins2, ins3 = program[pos + 1], program[pos + 2], program[pos + 3]
            if (ins1 if instructions[2] == 1 else program[ins1]) == (
                ins2 if instructions[1] == 1 else program[ins2]
            ):
                program[ins3] = 1
            else:
                program[ins3] = 0
            pos += 4
        else:
            assert opcode == 99
            break


v = 0
for sequences in phase_setting_sequences:
    output = 0
    for phase_setting in sequences:
        output = runner([phase_setting, output])
        print(phase_setting, output)
    if output > v:
        v = output
print(v)
