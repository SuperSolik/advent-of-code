import fileinput

from collections import deque
from itertools import pairwise

Vertex = tuple[int, int]
Graph = dict[Vertex, dict[Vertex, bool]]


def in_bounds(x, y, w, h) -> bool:
    return 0 <= x < w and 0 <= y < h


def set_up(graph: Graph, field: list[str], x: int, y: int, w: int, h: int):
    if in_bounds(x, y - 1, w, h) and field[y - 1][x] in '7F|':
        graph.setdefault((x, y), {})[(x, y - 1)] = True
        graph.setdefault((x, y - 1), {})[(x, y)] = True


def set_down(graph: Graph, field: list[str], x: int, y: int, w: int, h: int):
    if in_bounds(x, y + 1, w, h) and field[y + 1][x] in 'LJ|':
        graph.setdefault((x, y), {})[(x, y + 1)] = True
        graph.setdefault((x, y + 1), {})[(x, y)] = True


def set_right(graph: Graph, field: list[str], x: int, y: int, w: int, h: int):
    if in_bounds(x + 1, y, w, h) and field[y][x + 1] in '7J-':
        graph.setdefault((x, y), {})[(x + 1, y)] = True
        graph.setdefault((x + 1, y), {})[(x, y)] = True


def set_left(graph: Graph, field: list[str], x: int, y: int, w: int, h: int):
    if in_bounds(x - 1, y, w, h) and field[y][x - 1] in 'FL-':
        graph.setdefault((x, y), {})[(x - 1, y)] = True
        graph.setdefault((x - 1, y), {})[(x, y)] = True


def build_graph(lines: list[str]) -> tuple[Graph, Vertex]:
    graph = {}
    start = None

    h, w = len(lines), len(lines[0])
    for y in range(h):
        for x in range(w):
            match lines[y][x]:
                case '|':
                    set_down(graph, lines, x, y, w, h)
                    set_up(graph, lines, x, y, w, h)
                case '-':
                    set_left(graph, lines, x, y, w, h)
                    set_right(graph, lines, x, y, w, h)
                case 'L':
                    set_up(graph, lines, x, y, w, h)
                    set_right(graph, lines, x, y, w, h)
                case 'J':
                    set_up(graph, lines, x, y, w, h)
                    set_left(graph, lines, x, y, w, h)
                case '7':
                    set_left(graph, lines, x, y, w, h)
                    set_down(graph, lines, x, y, w, h)
                case 'F':
                    set_right(graph, lines, x, y, w, h)
                    set_down(graph, lines, x, y, w, h)
                case 'S':
                    set_up(graph, lines, x, y, w, h)
                    set_down(graph, lines, x, y, w, h)
                    set_right(graph, lines, x, y, w, h)
                    set_left(graph, lines, x, y, w, h)
                    start = (x, y)

    return graph, start


def find_loop(graph: Graph, start: Vertex) -> list[Vertex]:
    visited = set([start])
    stack = [start]
    parents = {start: None}
    cycles = []
    
    # dfs
    while stack:
        current_node = stack.pop()
        parent = parents.get(current_node)

        for neighbor in graph[current_node]:
            if neighbor not in visited:
                stack.append(neighbor)
                visited.add(neighbor)
                parents[neighbor] = current_node
            elif neighbor != parent:
                cycle_path = [neighbor, current_node]
                node = current_node

                while node != start:
                    node = parents.get(node)
                    cycle_path.append(node)

                cycles.append(cycle_path[::-1])

    return sorted(cycles, key=len)[-1]


def solve(lines: list[str]) -> None:
    graph, start = build_graph(lines)
    loop = find_loop(graph, start)

    after_start = loop[1]

    before_start = loop[-1]
    
    # check what pipe is on S
    start_c = None
    start_dir = (-after_start[0] + before_start[0], -after_start[1] + before_start[1])

    match start_dir:
        case (_, 0):
            start_c = '-'
        case (0, _):
            start_c = '|'
        case (-1, 1):
            start_c = 'F'
        case (-1, -1):
            start_c = 'L'
        case (1, 1):
            start_c = '7'
        case (1, -1):
            start_c = 'J'
    
    new_lines = []
    for y, line in enumerate(lines):
        new_lines.append(''.join(c if (x, y) in loop else '.' for x, c in enumerate(line)))

    lines = new_lines 
    lines[start[1]] = lines[start[1]].replace('S', start_c)

    outside = set()
    for y, row in enumerate(lines):
        inside = False
        pipe_start = None
        for x, c in enumerate(row):
            match c:
                case '|':
                    inside = not inside
                case 'L' | 'F':
                    pipe_start = c
                case '7' | 'J':
                    match pipe_start + c:
                        case 'LJ' | 'F7':
                            # movin along the pipe
                            pass
                        case _:
                            inside = not inside

                    pipe_start = None
                case _:
                    pass

            if not inside:
                outside.add((x, y))

    outside = outside.union(loop)
    print(len(lines) * len(lines[0]) - len(outside))


if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
