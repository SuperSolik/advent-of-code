import fileinput

from itertools import pairwise


def extrapolate(numbers: list[int]) -> int:
    steps = [numbers]
    while not all(x == 0 for x in steps[-1]):
        steps.append([b - a for a, b in pairwise(steps[-1])])

    return sum(s[-1] for s in steps)


def solve(lines: list[str]) -> None:
    print(
        sum(
            extrapolate([int(x) for x in line.split()])
            for line in lines
        )
    )
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
