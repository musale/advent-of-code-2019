ans = 0
start = 183564
stop = 657474
for pwd in range(start, stop + 1):
    digits = [int(x) for x in str(pwd)]
    has_dec = any([digits[i] > digits[i + 1] for i in range(len(digits) - 1)])
    has_pair = any(
        [
            (i == 0 or digits[i] != digits[i - 1])
            and digits[i] == digits[i + 1]
            and (i == len(digits) - 2 or digits[i] != digits[i + 2])
            for i in range(len(digits) - 1)
        ]
    )
    if has_pair and not has_dec:
        ans += 1
print(ans)
