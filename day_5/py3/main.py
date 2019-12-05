program = [int(i) for i in open("day_5/input.txt").read().split(",")]

pos = 0
while True:
    instructions = [int(x) for x in str(program[pos])]
    opcode = (0 if len(instructions) == 1 else instructions[-2]) * 10 + instructions[-1]
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
        program[ins1] = 1
        pos += 2
    elif opcode == 4:
        ins1 = program[pos + 1]
        print(program[ins1])
        pos += 2
    else:
        assert opcode == 99
        break
