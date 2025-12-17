#!/usr/bin/env python

import fileinput
from typing import Iterable


def solve1(lines: Iterable[str]) -> int:
    points = list(map(lambda line: tuple(map(int, line.split(","))), lines))

    max_area = 0
    for i in range(len(points)):
        for j in range(len(points)):
            if i != j:
                x1, y1 = points[i]
                x2, y2 = points[j]

                area = abs(x2 - x1 + 1) * abs(y2 - y1 + 1)
                if area > max_area:
                    max_area = area

    return max_area


def solve2(lines: Iterable[str]) -> None:
    pass


if __name__ == "__main__":
    data = [line.strip() for line in fileinput.input()]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
