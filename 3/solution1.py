#!/usr/bin/env python3

import fileinput


def solve(lines: list[str]) -> None:
    numbers = []
    result = 0

    cur_number = ''
    cur_cnt = 0

    h = len(lines)
    w = len(lines[0])

    for y in range(h):
        for x in range(w):
            c = lines[y][x]
            if not c.isdigit():
                # remember cur part number
                if cur_number and cur_cnt > 0:
                    numbers.append(cur_number)
                    result += int(cur_number)
                
                # reset
                cur_number = ''
                cur_cnt = 0
                continue

            cur_number += c

            # count neighbors
            for i in range(-1, 2):
                for j in range(-1, 2):
                    if 0 <= y + i < h and 0 <= x + j< w:
                        if lines[y + i][x + j] != '.' and not lines[y + i][x + j].isdigit():
                            cur_cnt += 1
        # end of line case
        # remember cur part number
        if cur_number and cur_cnt > 0:
            numbers.append(cur_number)
            result += int(cur_number)
        
        # reset
        cur_number = ''
        cur_cnt = 0

    print(result)

            

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])

