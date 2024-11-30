import fileinput
import heapq

from typing import Sequence, NamedTuple
from collections import deque


class Field:
    def __init__(self, points: list[Sequence[str]], w: int, h: int):
        self.points = points
        self.weights = {}
        self.w = w
        self.h = h

Coords = tuple[int, int]

def expand(field: Field, expand_k: int = 1) -> Field:
    field.h = len(field.points)
    field.w = len(field.points[0])

    # expand rows
    for y in range(field.h):
        if field.points[y].count('#') == 0:
            for x in range(field.w):
                field.weights[(x, y)] = expand_k
            

    # expand cols
    for x in range(field.w):
        col = [field.points[y][x] for y in range(field.h)]
        if col.count('#') == 0:
            for y in range(field.h):
                field.weights[(x, y)] = expand_k

    field.h = len(field.points)
    field.w = len(field.points[0])
    return field


def in_bounds(point: Coords, field: Field) -> bool:
    x, y = point
    return 0 <= x < field.w and 0 <= y < field.h


def dijkstra(field: Field, start: Coords, targets: list[Coords]) -> dict[Coords, int]:
    INF = 1e100
    dist = {start: 0}
    q = [(0, start)]

    while q:
        d, (x, y) = heapq.heappop(q)

        if all(t in dist for t in targets):
            break

        if d > dist.get((x, y), INF):
            continue

        for dx, dy in ((-1, 0), (1, 0), (0, 1), (0, -1)):
            neighbor = (x + dx, y + dy)
            new_dist = d + field.weights.get(neighbor, 1) # by default it's 1 step between coords, except for expanded
            if in_bounds(neighbor, field) and new_dist < dist.get(neighbor, INF):
                dist[neighbor] = new_dist
                heapq.heappush(q, (new_dist, neighbor))

    return dist


def solve(lines: list[str]) -> Field:
    field = Field([list(s) for s in lines], len(lines[0]), len(lines))

    # part 1 - expand x2
    # field = expand(field, 2)

    # part 2 - expand x1_000_000
    field = expand(field, 1_000_000)

    galaxies = []
    for y in range(field.h):
        for x in range(field.w):
            if field.points[y][x] == '#':
                galaxies.append((x, y))


    pairs = [(i, len(galaxies)) for i in range(len(galaxies))]
    result = 0
    for i, galaxy in enumerate(galaxies):
        start, end = pairs[i]
        other_galaxies = galaxies[start:end]
        dists = dijkstra(field, galaxy, other_galaxies)
        for o in other_galaxies:
            result += dists[o]
    
    print(result)

    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
