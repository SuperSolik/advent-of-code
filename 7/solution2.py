import fileinput

from functools import reduce
from collections import Counter


s = {
    'J': 1, # joker now has the lowest power
    '2': 2, 
    '3': 3, 
    '4': 4, 
    '5': 5, 
    '6': 6, 
    '7': 7, 
    '8': 8, 
    '9': 9, 
    'T': 10,
    'Q': 12,
    'K': 13,
    'A': 14
}


def resolve_joker(hand) -> str:
    j_cnt = hand.count('J')
    if j_cnt == 0 or j_cnt == 5:
        return hand
    
    # backtracking time
    # to see what is the best substitute for the jokers

    initial_cards = list(hand)
    non_jokers = list(hand.replace('J', ''))
    
    variants = set()

    def _resolve(cards: list[str], i: int):
        if i >= len(hand):
            nonlocal variants
            variants.add(''.join(cards))
            return 

        if cards[i] != 'J':
            return _resolve(cards, i+1)

        for c in non_jokers:
            new_cards = cards[:]
            new_cards[i] = c
            _resolve(new_cards, i+1)

    _resolve(initial_cards, 0)
    return sorted(variants, key=lambda x: get_hand_type(x))[-1]


def get_hand_type(hand: str) -> None:
    strongest = sorted(hand, key=lambda x: s[x])[-1]

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
    hands_data = []

    for line in lines:
        hand, bid = line.split()
        hands_data.append((hand, resolve_joker(hand), bid))
    
    # now we get the hand type using joker resolved version, but sort lexicographically still by original
    hands_data.sort(key=lambda hand_bid: (get_hand_type(hand_bid[1]), tuple(s[x] for x in hand_bid[0])))

    r = 0
    for i, (_, _, bid) in enumerate(hands_data):
        r += (i + 1) * int(bid)
    
    print(r)
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
