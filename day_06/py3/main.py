from collections import deque

map_values = [v.strip().split(")") for v in open("day_06/input.txt").readlines()]

values = {}
for value in map_values:
    a, b = value
    if a not in values:
        values[a] = []
    if b not in values:
        values[b] = []
    values[a].append(b)
    values[b].append(a)

distance = {}
q = deque()
q.append(("YOU", 0))

while q:
    a, b = q.popleft()
    if a in distance:
        continue
    distance[a] = b
    for i in values[a]:
        q.append((i, b + 1))

print(distance["SAN"] - 2)


def loop(k):
    s = 0
    for i in values.get(k, []):
        s += loop(i)
        s += 1
    return s


# total = 0
# for k in values:
#     total += loop(k)

# print(total)
