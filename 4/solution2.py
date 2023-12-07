import fileinput


cards_count = {}


def process_card(card_idx: int, line: str) -> int:
    numbers = line.split(':')[1].strip()
    winning, have = numbers.split('|')
    winning = winning.split()
    have = have.split()
    

    have_winning = 0
    for n in have:
        if n in winning:
            have_winning += 1

    cur_card_count = cards_count.get(card_idx, 1)

    for i in range(have_winning):
        nxt_card_count = cards_count.get(card_idx + i + 1, 1)
        cards_count[card_idx + i + 1] = nxt_card_count + cur_card_count

    return cur_card_count


def solve(lines: list[str]) -> None:
    print(sum(process_card(i, line) for i, line in enumerate(lines)))
            

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
