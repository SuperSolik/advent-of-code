import fileinput
import re

from collections import deque

ins_pat = re.compile(r'^([LRDU])\s(\d+)\s\((#[\w]{6})\)')

def flood_fill(field: list[list[str]], w: int, h: int, fill_val: int) -> list[list[str]]:
    start = 0, 0
    q = deque([start])

    while q:
        x, y = q.pop()
        
        if field[y][x] > 0:
            continue

        field[y][x] = fill_val

        for (dx, dy) in ((-1, 0), (1, 0), (0, -1), (0, 1)):
            if 0 <= x + dx < w and 0 <= y + dy < h:
                q.appendleft((x + dx, y + dy))

    return field


def solve(lines: list[str]) -> None:
    size = 800
    field = [[0 for _ in range(size)] for _ in range(size)]

    x, y = size // 2, size // 2
    field[y][x] = 1

    for line in lines:
        direction, cnt, color = ins_pat.match(line).groups()
        cnt = int(cnt)
        
        dd = (None, None)
        match direction:
            case 'R':
                dd = (1, 0)
            case 'L':
                dd = (-1, 0)
            case 'D':
                dd = (0, 1)
            case 'U':
                dd = (0, -1)

        for i in range(cnt):
            x += dd[0]
            y += dd[1]
            field[y][x] = 1
    
    fill_val = 2
    field = flood_fill(field, size, size, fill_val)   

    res = 0
    for row in field:
        res += (size - row.count(fill_val))

    print(res)
                
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
