import fileinput

from typing import Sequence
from copy import deepcopy


def sift(field: list[list[str]], w: int, h: int) -> list[list[str]]:
    for x in range(w):
        for _ in range(h):
            for y in range(h):
                if field[y][x] in '.#':
                    continue

                if y > 0 and field[y-1][x] not in '#O':
                    field[y-1][x] = 'O' 
                    field[y][x] = '.'
    return field

def rotate(lines: list[list[str]]) -> list[list[str]]:
    return list(map(list, list(zip(*lines[::-1]))))


def score(lines: Sequence[Sequence[str]]) -> int:
    res = 0
    for j, line in enumerate(lines):
        res += (len(lines) - j) * line.count('O')
    return res


def solve(lines: list[str]) -> None:
    lines = list(map(list, lines))
    
    h, w = len(lines), len(lines[0])
    

    def cycle():
        nonlocal lines, w, h
        for _ in range(4):
            lines = sift(lines, w, h)
            w, h = h, w
            lines = rotate(lines)

    done = set()
    cycle_results = [] 

    i = 0
    def froze(lines: list[list[str]]) -> tuple[tuple[str]]:
        return tuple(tuple(l) for l in lines)

    seen = set()
    variants = []

    limit = 1000000000
    while i < limit:

        cycle()

        f = froze(lines)

        if f in seen:
            break

        i += 1
        seen.add(f)
        variants.append(f)


    # so this is roughly how it looks
    # 0 .... x at i ..<j - i steps>.. x at j ....... limit - 1
    #                                          ^^^ cycle repeats here

    # calc which idx we'll end up on
    fst = variants.index(f)
    limit -= fst
    idx = limit % (i - fst) + fst

    print(score(variants[idx-1]))

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
