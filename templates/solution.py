#!/usr/bin/env python

import fileinput
from typing import Iterable


def solve1(lines: Iterable[str]) -> None:
    pass


def solve2(lines: Iterable[str]) -> None:
    pass


if __name__ == "__main__":
    data = [line.strip() for line in fileinput.input()]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
