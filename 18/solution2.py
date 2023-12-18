import fileinput
import re

from itertools import pairwise


ins_pat = re.compile(r'^([LRDU])\s(\d+)\s\(#([\w]{6})\)')
dirs = {0: 'R', 1: 'D', 2: 'L', 3: 'U'}


def solve(lines: list[str]) -> None:
    points = [(0, 0)]
    border = 0

    for line in lines:
        x, y = points[-1]
        _, _, color = ins_pat.match(line).groups()

        cnt, direction = color[:-1], color[-1]

        cnt = int(cnt, 16)
        direction = int(direction, 16)


        cnt = int(cnt)
        
        dd = (None, None)
        match dirs[direction]:
            case 'R':
                dd = (1, 0)
            case 'L':
                dd = (-1, 0)
            case 'D':
                dd = (0, 1)
            case 'U':
                dd = (0, -1)

        x += dd[0] * cnt
        y += dd[1] * cnt
        border += cnt

        points.append((x, y))

    # learning something new here:
    
    # step 1: Shoelace theorem
    # https://en.wikipedia.org/wiki/Shoelace_formula

    area = 0
    for i in range(len(points)):
        x = points[i][0]
        y_prev = points[i-1][1]
        y_next = points[(i+1) % len(points)][1]

        area += x * (y_next - y_prev)

    area = abs(area) // 2


    # step 2: Pick's theorem
    # https://en.wikipedia.org/wiki/Pick%27s_theorem

    # area = inner + border / 2 - 1, we need the inner, so inner = area - border / 2 + 1
    inner = area - border // 2 + 1

    print(inner + border)
    
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
