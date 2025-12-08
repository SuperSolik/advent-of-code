#!/usr/bin/env python

import fileinput
import heapq

from typing import Iterable
from math import inf


def solve1(lines: Iterable[str]) -> int:
    points = []
    for line in lines:
        x, y, z = tuple(map(int, line.split(",")))
        points.append((x, y, z))

    p_len = len(points)

    dists = [inf] * p_len * p_len

    # NOTE: calc distance between all points
    for i in range(p_len):
        for j in range(p_len):
            (x1, y1, z1) = points[i]
            (x2, y2, z2) = points[j]
            dist = (x1 - x2) ** 2 + (y1 - y2) ** 2 + (z1 - z2) ** 2

            dists[i * p_len + j] = dist
            dists[j * p_len + i] = dist

    dists = [(dist, i) for i, dist in enumerate(dists)]

    # NOTE: order pairs of points by distance via heapq
    heapq.heapify(dists)
    p = heapq.heappop(dists)

    # NOTE: skip self connections (point with itself)
    # need to do pop twice as each dist is duplicated (i -> j, j -> i)
    while p[0] == 0:
        heapq.heappop(dists)
        p = heapq.heappop(dists)

    circuits = [False] * p_len * p_len

    # NOTE: form graph
    for i in range(999):
        idx = p[1]
        i = idx // p_len
        j = idx % p_len

        circuits[i * p_len + j] = True
        circuits[j * p_len + i] = True

        heapq.heappop(dists)
        p = heapq.heappop(dists)

    visited = [False] * p_len

    def dfs(node: int, component: list[int]):
        visited[node] = True
        component.append(node)
        for i, n in enumerate(circuits[node * p_len : (node + 1) * p_len]):
            if n and not visited[i]:
                dfs(i, component)

    components = []
    for node in range(p_len):
        if not visited[node]:
            component = []
            dfs(node, component)
            components.append(component)

    a, b, c = sorted(map(lambda c: len(c), components), reverse=True)[:3]
    return a * b * c


def solve2(lines: Iterable[str]) -> None:
    points = []
    for line in lines:
        x, y, z = tuple(map(int, line.split(",")))
        points.append((x, y, z))

    p_len = len(points)

    dists = [inf] * p_len * p_len

    for i in range(p_len):
        for j in range(p_len):
            (x1, y1, z1) = points[i]
            (x2, y2, z2) = points[j]
            dist = (x1 - x2) ** 2 + (y1 - y2) ** 2 + (z1 - z2) ** 2

            dists[i * p_len + j] = dist
            dists[j * p_len + i] = dist

    dists = [(dist, i) for i, dist in enumerate(dists)]

    circuits = [False] * p_len * p_len

    heapq.heapify(dists)
    p = heapq.heappop(dists)

    while p[0] == 0:
        heapq.heappop(dists)
        p = heapq.heappop(dists)

    # NOTE: same as part 1, except do full connection test
    while True:
        idx = p[1]
        i = idx // p_len
        j = idx % p_len

        circuits[i * p_len + j] = True
        circuits[j * p_len + i] = True

        # NOTE: full connection means we have at least one connection in each row of adj matrix
        if all(any(circuits[y * p_len : (y + 1) * p_len]) for y in range(p_len)):
            break

        heapq.heappop(dists)
        p = heapq.heappop(dists)

    return points[i][0] * points[j][0]


if __name__ == "__main__":
    data = [line.strip() for line in fileinput.input()]
    print("Part1: ", solve1(data))
    print("Part2: ", solve2(data))
