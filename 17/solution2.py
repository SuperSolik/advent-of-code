import fileinput

from heapq import heappop, heappush


def solve(lines: list[str]) -> None:
    w, h = len(lines[0]), len(lines)
    

    q = [(0, (0, 0), (0, 0), 0)]
    visited = set()

    while q:
        node = heappop(q)
        (hl, (x, y), (dx, dy), n) = node

        if x == w - 1 and y == h - 1 and n >= 4:
            break

        if (x, y, dx, dy, n) in visited:
            continue

        visited.add((x, y, dx, dy, n))

        if n < 10 and (dx, dy) != 0:
            if 0 <= x + dx < w and 0 <= y + dy < h:
              heappush(q, (hl + int(lines[y + dy][x + dx]), (x + dx, y + dy), (dx, dy), n + 1))

        if n >= 4 or (dx, dy) == (0, 0):
            for (ndx, ndy) in set([(-1, 0), (1, 0), (0, -1), (0, 1)]).difference([(dx, dy), (-dx, -dy)]):
                if 0 <= x + ndx < w and 0 <= y + ndy < h:
                    heappush(q, (hl + int(lines[y + ndy][x + ndx]), (x + ndx, y + ndy), (ndx, ndy), 1))

    print(hl)
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
