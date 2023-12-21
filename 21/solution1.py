import fileinput
import heapq


def solve(lines: list[str]) -> None:
    w, h = len(lines[0]), len(lines)

    start = (None, None)
    for y in range(h):
        for x in range(w):
            if lines[y][x] == 'S':
                start = (x, y)
                break

    INF = 1e100
    dist = {start: 0}
    q = [(0, start)]


    MAX_STEPS = 64
    
    while q:
        d, (x, y) = heapq.heappop(q)

        if d > dist.get((x, y), INF):
            continue

        for dx, dy in ((-1, 0), (1, 0), (0, 1), (0, -1)):
            neighbor = (x + dx, y + dy)
            new_dist = d + 1
            if 0 <= x + dx < w and 0 <= y + dy < h and lines[y + dy][x + dx] in '.S' and new_dist < dist.get(neighbor, INF):
                dist[neighbor] = new_dist
                heapq.heappush(q, (new_dist, neighbor))
        
    i = 0
    for c, d  in dist.items():
        if d <= MAX_STEPS and d % 2 == 0:
            i += 1
    
    print(i)

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
