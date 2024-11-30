import fileinput


def solve(lines: list[str]) -> None:
    lines = list(map(list, lines))
    
    h, w = len(lines), len(lines[0])
    for x in range(w):
        for _ in range(h):
            for y in range(h):
                if lines[y][x] in '.#':
                    continue

                if y > 0 and lines[y-1][x] not in '#O':
                    lines[y-1][x] = 'O' 
                    lines[y][x] = '.'
    
    res = 0
    for i, line in enumerate(lines):
        res += (h - i) * line.count('O')

    print(res)
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
