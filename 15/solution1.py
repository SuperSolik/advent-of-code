import fileinput

def xmas_hash(s: str) -> int:
    h = 0
    for ch in s:
        h += ord(ch)
        h *= 17
        h = h % 256
    return h

def solve(lines: list[str]) -> None:
    res = 0
    for p in lines[0].split(','):
        res += xmas_hash(p)

    print(res)
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
