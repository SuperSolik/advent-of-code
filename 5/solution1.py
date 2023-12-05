#!/usr/bin/env python3

import fileinput
import re

from math import inf
from pprint import pprint


class RangeMap(dict):
    def get_range(self, x):
        prev = -1000_000_000

        # sort by start of range
        for r in sorted(self.keys(), key=lambda r: r[0]):
            start, end = r

            if prev <= x < start:
                # if x is bigger then prev range, but lesser then next range
                # -> because we sorted, we know then there won't be matching inverval further
                # so we can return early
                # 
                # this may be buggy, can be deleted
                return None

            if start <= x < end:
                return self.__getitem__(r), r

            prev = end

        return None


def parse_map(lines: list[str]) -> tuple[list[str], tuple[str, RangeMap]]:
    heading = lines[0]
    name, _ = heading.split()

    i = 1
    map_data = RangeMap()
    while i < len(lines):
        if lines[i] == '\n':
            # end of map
            break

        range_data = lines[i]
        dest_start, source_start, step = [int(x) for x in range_data.split()]

        map_data[(source_start, source_start + step)] = (dest_start, dest_start + step) 
            
        i += 1
        
    return lines[i+1:], (name, map_data)

def solve(lines: list[str]) -> None:
    seeds, lines = lines[0], lines[2:]
    seeds = [int(s) for s in seeds.split(':')[1].strip().split()]

    maps = {}
    while lines:
        lines, (map_name, map_data) = parse_map(lines)
        maps[map_name] = map_data
    
    input_ = seeds
    for transformation in maps.keys():
        for i in range(len(seeds)):
            res = maps[transformation].get_range(seeds[i])
            if res is not None:
                (target_start, _), (source_start, _) = res
                seeds[i] = target_start + (seeds[i] - source_start)
            
    print(min(seeds))

if __name__ == '__main__':
    solve([line for line in fileinput.input()])

