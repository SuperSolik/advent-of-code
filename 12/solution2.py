import fileinput
import re

from functools import cache


@cache
def fit_cnt(pat: str, groups: tuple[int]) -> int:
    if not groups:
        if '#' in pat:
            return 0
        return 1

    if not pat:
        return 0

 
    match pat[0]:
        case '.':
            return fit_cnt(pat[1:], groups)
        case '#':
            g = groups[0]
            place = pat[:g]
            if place.replace('?', '#') != '#' * g:
                return 0

            if len(pat) == g:
                if len(groups) == 1:
                    return 1
                return 0

            if pat[g] in '?.':
                return fit_cnt(pat[g+1:], groups[1:])
            
            return 0
        case '?':
            return fit_cnt('.' + pat[1:], groups) + fit_cnt('#' + pat[1:], groups)
        case _:
            raise RuntimeError('unreachable')


def solve(lines: list[str]) -> None:
    result = 0

    for line in lines:
        pattern, groups = line.split()
        groups = tuple(int(x) for x in groups.split(','))
        result += fit_cnt('?'.join([pattern] * 5), groups*5)

    print(result)
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
