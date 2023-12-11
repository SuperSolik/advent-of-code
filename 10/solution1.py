import fileinput

from collections import deque

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


def calc_max_dist(graph: Graph, lines: list[str], start: Vertex) -> int:
    dist = {start: 0}
    visited = set()

    stack = [start]
    q = deque([start])
    
    while q:
        cur = q.popleft()
        
        if cur not in visited:
            visited.add(cur)
            for v in graph[cur].keys():
                if v not in visited:
                    dist[v] = max(dist.get(v, 0), dist[cur] + 1)
                else:
                    dist[v] = min(dist.get(v, 0), dist[cur] + 1)

                q.append(v)

    return max(dist.values())

        


def solve(lines: list[str]) -> None:
    graph, start = build_graph(lines)
    print(calc_max_dist(graph, lines, start))
    

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
