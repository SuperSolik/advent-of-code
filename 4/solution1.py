#!/usr/bin/env python3

import fileinput


def parse_card(line: str) -> int:
    numbers = line.split(':')[1].strip()
    winning, have = numbers.split('|')
    winning = winning.split()
    have = have.split()
    

    card_points = 0
    for n in have:
        if n in winning:
            card_points = card_points * 2 if card_points else 1


    return card_points


def solve(lines: list[str]) -> None:
    print(sum(parse_card(line) for line in lines))
            

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])

