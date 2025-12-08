#!/usr/bin/env python

import fileinput
from typing import Sequence
from functools import lru_cache


def solve1(rows: Sequence[str]) -> int:
    result = 0
    beams = set((rows[0].index("S"),))

    for row in rows[1:]:
        for x, c in enumerate(row):
            if c == "^" and x in beams:
                result += 1
                # NOTE: beam hit the splitter, split
                beams.remove(x)
                beams.add(x - 1)
                beams.add(x + 1)

    return result


def solve2(rows: Sequence[str]) -> int:
    # NOTE: too slow due to recursion and doesn't produce the output because of that, need to work on this more

    beam_path = (rows[0].index("S"),)
    timelines = set()
    ll = len(rows)

    @lru_cache(maxsize=None)
    def step(beam_path: tuple[int], sy: int):
        if sy == ll - 1:
            # NOTE: reached the end
            timelines.add(tuple(beam_path))
            return

        has_splitters = False
        for x, c in enumerate(rows[sy]):
            if c == "^" and x == beam_path[-1]:
                step(beam_path + (x - 1,), sy + 1)
                step(beam_path + (x + 1,), sy + 1)
                has_splitters = True

        if not has_splitters:
            step(beam_path, sy + 1)

    step(beam_path, 1)

    return len(timelines)


if __name__ == "__main__":
    data = [line.strip() for line in fileinput.input()]
    print("Part1: ", solve1(data))
    # TODO: part2
    # print("Part2: ", solve2(data))
