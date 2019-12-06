map_values = [v.strip().split(")") for v in open("day_6/input.txt").readlines()]

values = {}
for value in map_values:
    a, b = value
    if a not in values:
        values[a] = []
    if b not in values:
        values[b] = []
    values[a].append(b)
    # values[b].append(a)


def loop(k):
    s = 0
    for i in values.get(k, []):
        s += loop(i)
        s += 1
    return s


total = 0
for k in values:
    total += loop(k)

print(total)
