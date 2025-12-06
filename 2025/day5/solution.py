#!/usr/bin/env python

import fileinput

IntervalList = list[tuple[int, int]]


def parse_input(
    input: fileinput.FileInput[str],
) -> tuple[IntervalList, list[int]]:
    intervals = []
    products = []

    parsing_products = False

    for line in input:
        line = line.strip()

        if not line:
            parsing_products = True
            continue

        if parsing_products:
            products.append(int(line))
        else:
            intervals.append(tuple(map(int, line.split("-"))))

    return intervals, products


def merge_intervals(intervals: IntervalList) -> IntervalList:
    merged: IntervalList = []

    for interval in sorted(intervals, key=lambda i: i[0]):
        # NOTE: if we're inserting first interval
        # or if the currently merged one doesn't overlap with the current interval
        # intervals are sorted by start, so we only have to check if the end of the merged overlaps with the start of the current
        if not merged or not merged[-1][1] >= interval[0]:
            merged.append(interval)
        else:
            # NOTE: merge = take the min start and max end
            merged[-1] = (merged[-1][0], max(merged[-1][1], interval[1]))

    return merged


def solve1(intervals: IntervalList, products: list[int]) -> int:
    result = 0

    for p in products:
        if p < intervals[0][0] or p > intervals[-1][1]:
            continue

        for i in intervals:
            if i[0] <= p <= i[1]:
                result += 1

    return result


def solve2(intervals: IntervalList, _: list[int]) -> int:
    return sum(i[1] - i[0] + 1 for i in intervals)


if __name__ == "__main__":
    intervals, products = parse_input(fileinput.input())
    merged_intervals = merge_intervals(intervals)

    print("Part1: ", solve1(merged_intervals, products))
    print("Part2: ", solve2(merged_intervals, products))
