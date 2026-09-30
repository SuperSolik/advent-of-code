#!/usr/bin/env python

import fileinput
from collections.abc import Iterable


def solve1(lines: Iterable[str]) -> None:
    print(lines)


def solve2(lines: Iterable[str]) -> None:
    print(lines)


if __name__ == "__main__":
    with fileinput.input(encoding="utf-8") as stdin:
        lines = [line.strip() for line in stdin]

    print("Part1: ", solve1(lines))
    print("Part2: ", solve2(lines))
