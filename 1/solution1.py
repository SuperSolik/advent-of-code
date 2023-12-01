import fileinput


def parse_line(line: str) -> int:
    pair = [None, None]
    for i, c in enumerate(line):
        if not c.isdigit():
            continue
        
        if pair[0] is None:
            pair[0] = int(c)
        else:
            pair[1] = int(c)

    if pair[1] is None:
        pair[1] = pair[0]

    return pair[0] * 10 + pair[1] if all(pair) else 0
    

if __name__ == '__main__':
    print(sum((parse_line(line) for line in fileinput.input())))
