import fileinput


def solve(lines: list[str]) -> None:
    print(lines)
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
