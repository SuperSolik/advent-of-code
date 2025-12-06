#!/usr/bin/env python

import fileinput
from functools import reduce
from operator import add, mul
from typing import Iterable


def solve1(lines: Iterable[str]) -> int:
    result = 0

    for el in zip(*(line.strip().split() for line in lines)):
        numbers = tuple(map(int, el[:-1]))
        op = add if el[-1] == "+" else mul

        result += reduce(lambda x, y: op(x, y), numbers)

    return result


def solve2(lines: Iterable[str]) -> None:
    result = 0
    numbers = []
    op = None

    for el in zip(*lines):
        el = "".join(el)

        if not el.strip():
            result += reduce(lambda x, y: op(x, y), numbers)
            numbers = []
            op = None
            continue

        number_part = el[:-1]

        numbers.append(int(number_part.strip()))
        if el[-1] != " ":
            op = add if el[-1] == "+" else mul

    result += reduce(lambda x, y: op(x, y), numbers)
    return result


if __name__ == "__main__":
    data = [line for line in fileinput.input()]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
