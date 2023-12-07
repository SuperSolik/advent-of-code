#!/usr/bin/env python3

import fileinput

from functools import reduce
from collections import Counter


s = {
    '2': 2, 
    '3': 3, 
    '4': 4, 
    '5': 5, 
    '6': 6, 
    '7': 7, 
    '8': 8, 
    '9': 9, 
    'T': 10,
    'J': 11,
    'Q': 12,
    'K': 13,
    'A': 14
}


def get_hand_type(hand: str) -> None:
    hand_cnt = Counter(hand).most_common(5)
    match hand_cnt:
        case [(x, 5)]:
            return 5
        case [(x, 4), (y, 1)]:
            return 4
        case [(x, 3), (y, 2)]:
            return 3
        case [(x, 3), (y, 1), (z, 1)]:
            return 2
        case [(x, 2), (y, 2), (z, 1)]:
            return 1
        case [(x, 2), (y, 1), (z, 1), (w, 1)]:
            return 0
        case [(x, 1), (y, 1), (z, 1), (w, 1), (q, 1)]:
            return -1
        case _:
            raise ValueError(f'terrible hand of {hand_cnt}')


def solve(lines: list[str]) -> None:
    hands_data = [line.split() for line in lines]

    hands_data.sort(key=lambda hand_bid: (get_hand_type(hand_bid[0]), tuple(s[x] for x in hand_bid[0])))

    r = 0
    for i, (hand, bid) in enumerate(hands_data):
        r += (i + 1) * int(bid)
    
    print(r)
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
