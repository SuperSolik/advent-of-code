import fileinput
import re
from typing import Iterable


def solve1(lines: Iterable[str]) -> int:
    result = 0
    for pair in lines:
        start, end = map(int, pair.split("-"))
        for num in range(start, end + 1):
            repr_ = str(num)
            len_num = len(repr_)
            if len_num % 2 != 0:
                continue
            part1 = repr_[: len_num // 2]
            part2 = repr_[len_num // 2 :]
            if part1 == part2:
                result += num

    return result


def solve2(lines: Iterable[str]) -> int:
    result = 0
    for pair in lines:
        start, end = map(int, pair.split("-"))
        for num in range(start, end + 1):
            repr_ = str(num)
            match = re.fullmatch(r"(\d+?)\1+", repr_)
            if match:
                result += num

    return result


if __name__ == "__main__":
    data = [
        pair for line in fileinput.input() for pair in line.strip().split(",") if pair
    ]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
