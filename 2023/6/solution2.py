import fileinput
import math


def solve_race(t, d) -> int:
    # find roots of t * x - x * x - d = 0

    x1 = (t - math.sqrt(t * t - 4 * d)) / 2
    x2 = (t + math.sqrt(t * t - 4 * d)) / 2
     
    r = math.ceil(x2 - 1) - math.floor(x1 + 1) + 1
    return r
    


def solve(lines: list[str]) -> None:
    race_time = int(''.join([t for t in lines[0].split()[1:]]))
    race_distance = int(''.join([d for d in lines[1].split()[1:]]))

    print(solve_race(race_time, race_distance)) 
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
