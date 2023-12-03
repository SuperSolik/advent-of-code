#!/usr/bin/env python3

import fileinput

from functools import reduce


def solve(lines: list[str]) -> None:
    coords_to_numbers = {}
    
    gears = []

    result = 0

    cur_number = ''
    cur_cnt = 0

    start_x = -1
    stary_y = -1

    h = len(lines)
    w = len(lines[0])

    for y in range(h):
        for x in range(w):
            c = lines[y][x]
            if not c.isdigit():
                if c == '*':
                    # save gear
                    gears.append((x, y))
                
                # remember cur part number coords
                if cur_number and cur_cnt > 0:
                    for num_x in range(start_x, start_x + len(cur_number)):
                        coords_to_numbers.setdefault(y, {})[num_x] = int(cur_number)

                cur_number = ''
                cur_cnt = 0
                start_x, start_y = -1, -1
                continue
            
            if not cur_number:
                # remember of the number
                start_x, start_y = x, y

            cur_number += c

            # count neighbors
            for i in range(-1, 2):
                for j in range(-1, 2):
                    if 0 <= y + i < h and 0 <= x + j< w:
                        if lines[y + i][x + j] != '.' and not lines[y + i][x + j].isdigit():
                            cur_cnt += 1
         
        # end of line case
        # remember cur part number coords
        if cur_number and cur_cnt > 0:
            for num_x in range(start_x, start_x + len(cur_number)):
                coords_to_numbers.setdefault(y, {})[num_x] = int(cur_number)

        cur_number = ''
        cur_cnt = 0
        start_x, start_y = -1, -1
    
    # check gears
    for gx, gy in gears:
        gear_nums = set()
        for i in range(-1, 2):
            for j in range(-1, 2):
                if 0 <= gy + i < h and 0 <= gx + j< w:
                    num = coords_to_numbers.get(gy + i, {}).get(gx + j)
                    if num:
                        gear_nums.add(num)

        if len(gear_nums) == 2:
            result += reduce(lambda x,y: x * y, gear_nums)

    print(result)

            

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])

