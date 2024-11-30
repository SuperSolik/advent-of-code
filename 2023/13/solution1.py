import fileinput

def check(lines: list[str]) -> int:
    for i in range(1, len(lines)):
        a, b = lines[:i][::-1], lines[i:]
        if a[:len(b)] == b[:len(a)]:
            return i

    return 0

def solve(lines: list[str]) -> None:
    patterns = []
    pattern = []
    for line in lines:
        if line == "":
            if pattern:
                patterns.append(pattern)
                pattern = []
        else:
            pattern.append(line)

    if pattern:
        patterns.append(pattern)

    res = 0

    for p in patterns:
        h = check(p)
        if h: 
            res += h * 100
        else:
            v = check(list(zip(*p)))
            res += v

    print(res)

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
