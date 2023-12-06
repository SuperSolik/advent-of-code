#!/usr/bin/env python3

import fileinput
import math

from functools import reduce


def solve_race(t, d) -> int:
    # find roots of t * x - x * x - d = 0

    x1 = (t - math.sqrt(t * t - 4 * d)) / 2
    x2 = (t + math.sqrt(t * t - 4 * d)) / 2
     
    r = math.ceil(x2 - 1) - math.floor(x1 + 1) + 1
    return r
    


def solve(lines: list[str]) -> None:
    times = [int(t) for t in lines[0].split()[1:]]
    distances = [int(d) for d in lines[1].split()[1:]]

    races = list(zip(times, distances))
    
    print(reduce(lambda x, y: x * y, (solve_race(*r) for r in races)))
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
