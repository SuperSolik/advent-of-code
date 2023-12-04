#!/usr/bin/env python3

import fileinput


cards_cnt = {}


def parse_card(card_idx: int, line: str) -> int:
    numbers = line.split(':')[1].strip()
    winning, have = numbers.split('|')
    winning = winning.split()
    have = have.split()
    

    have_winning = 0
    for n in have:
        if n in winning:
            have_winning += 1

    cur_card_cnt = cards_cnt.get(card_idx, 1)

    for i in range(have_winning):
        cur_card_count = cards_cnt.get(card_idx + i + 1, 1)
        cards_cnt[card_idx + i + 1] = cur_card_count + cur_card_cnt

    return cur_card_cnt


def solve(lines: list[str]) -> None:
    print(sum(parse_card(i, line) for i, line in enumerate(lines)))
            

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
