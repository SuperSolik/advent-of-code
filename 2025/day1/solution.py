import fileinput
from typing import Iterable


def solve1(lines: Iterable[str]) -> int:
    result = 0
    start = 50
    for line in lines:
        steps = int(line[1:])
        if line[0] == "L":
            steps = -steps

        start = (start + steps) % 100

        if start == 0:
            result += 1

    return result


def solve2(lines: Iterable[str]) -> int:
    result = 0
    start = 50
    for line in lines:
        steps = int(line[1:])
        if line[0] == "L":
            steps = -steps

        result += abs((start + steps) // 100)
        start = abs((start + steps) % 100)

    return result


if __name__ == "__main__":
    data = [line.strip() for line in fileinput.input()]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
