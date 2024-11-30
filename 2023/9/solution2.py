import fileinput

from itertools import pairwise
from functools import reduce


def extrapolate(numbers: list[int]) -> int:
    steps = [numbers]
    while not all(x == 0 for x in steps[-1]):
        steps.append([b - a for a, b in pairwise(steps[-1])])
    
    return reduce(lambda acc, x: x - acc, (s[0] for s in reversed(steps)))


def solve(lines: list[str]) -> None:
    print(
        sum(
            extrapolate([int(x) for x in line.split()])
            for line in lines
        )
    )
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
