import fileinput
from copy import deepcopy


def remove_rolls(space: list[list[str]]) -> int:
    result = 0
    to_remove = []
    w, h = len(space[0]), len(space)
    for y in range(0, h):
        for x in range(0, w):
            if space[y][x] != "@":
                continue

            neighbors_cnt = 0

            for dy in (-1, 0, 1):
                for dx in (-1, 0, 1):
                    if dx == 0 and dy == 0:
                        continue
                    if (
                        0 <= x + dx < w
                        and 0 <= y + dy < h
                        and space[y + dy][x + dx] == "@"
                    ):
                        neighbors_cnt += 1

            if neighbors_cnt < 4:
                result += 1
                to_remove.append((x, y))

    for x, y in to_remove:
        space[y][x] = "."

    return result


def solve1(lines: list[list[str]]) -> int:
    return remove_rolls(lines)


def solve2(lines: list[list[str]]) -> int:
    result = 0
    while True:
        removed_cnt = remove_rolls(lines)
        if removed_cnt == 0:
            break

        result += removed_cnt

    return result


if __name__ == "__main__":
    data = [list(line.strip()) for line in fileinput.input()]
    print("Part1: ", solve1(deepcopy(data)))
    print("Part2: ", solve2(deepcopy(data)))
