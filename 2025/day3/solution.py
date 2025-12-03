import fileinput
from typing import Iterable, Sequence


def _max(array: Sequence[int], start: int, end: int) -> tuple[int, int]:
    max_, idx = -1, -1
    for i in range(start, end):
        if array[i] > max_:
            max_, idx = array[i], i

    return max_, idx


def solve1(lines: Iterable[str]) -> int:
    result = 0
    for line in lines:
        array = tuple(map(int, line))
        l_a = len(array)
        a, idx_a = _max(array, 0, l_a - 1)
        b, _ = _max(array, idx_a + 1, l_a)

        result += a * 10 + b

    return result


def solve2(lines: Iterable[str]) -> int:
    result = 0
    for line in lines:
        array = tuple(map(int, line))
        l_a = len(array)
        max_num = []
        start = 0

        for i in range(0, 12):
            val, start = _max(array, start, l_a - 11 + i)
            max_num.append(val)
            start += 1

        result += int("".join(map(str, max_num)))

    return result


if __name__ == "__main__":
    data = [line.strip() for line in fileinput.input()]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
