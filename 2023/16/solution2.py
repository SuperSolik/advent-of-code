import fileinput

def try_beams(lines: list[str], w: int, h: int, start: tuple[int, int], direction: tuple[int, int]) -> int:
    energy = [[0 for _ in range(w)] for _ in range(h)]

    # first beam, from 0,0 moving to right
    sx, sy = start
    dx, dy = direction

    beams = [{
        "x": sx,
        "y": sy,
        "dx": dx,
        "dy": dy,
        "alive": True,
    }]

    def in_bounds(x: int, y: int) -> bool:
        nonlocal w, h
        return 0 <= x < w and 0 <= y < h

    visited = set()

    while beams:
        spawn_beams = []
        for i in range(len(beams)):
            x, y, dx, dy = beams[i]["x"], beams[i]["y"], beams[i]["dx"], beams[i]["dy"]

            if (x, y, dx, dy) in visited:
                # we've been there before
                beams[i]["alive"] = False
                continue

            # energize
            energy[y][x] = 1
            visited.add((x, y, dx, dy))

            spawn_opposite = False
            ndx, ndy = dx, dy

            match lines[y][x], (dx, dy):
                case '-', (0, _):
                    ndx, ndy = ndy, ndx
                    spawn_opposite = True
                case '|', (_, 0):
                    ndx, ndy = ndy, ndx
                    spawn_opposite = True
                case '\\', _:
                    ndx, ndy = dy, dx
                case '/', _:
                    ndx, ndy = -dy, -dx
                case _, _:
                    pass

            if spawn_opposite:
                spawn_beams.append({
                    "x": x,
                    "y": y,
                    "dx": -ndx,
                    "dy": -ndy,
                    "alive": True,
                })

            beams[i]["dx"] = ndx
            beams[i]["dy"] = ndy

            if in_bounds(x + beams[i]["dx"], y + beams[i]["dy"]):
                beams[i]["x"] = x + beams[i]["dx"]
                beams[i]["y"] = y + beams[i]["dy"]
            else:
                beams[i]["alive"] = False
        
        beams = [b for b in beams if b["alive"]]
        beams.extend(spawn_beams)

    return sum(e.count(1) for e in energy)


def solve(lines: list[str]) -> None:
    h, w = len(lines), len(lines[0])

    # list of [(sx, sy, dx, dy)]
    start_pos = []

    res = 0
    for i in range(h):
        start_pos.append((i, 0, 0, 1))
        start_pos.append((i, h - 1, 0, -1))
        start_pos.append((0, i, 1, 0))
        start_pos.append((h - 1, i, -1, 0))

    for sx, sy, dx, dy in start_pos:
        beam_res = try_beams(lines, w, h, (sx, sy), (dx, dy))
        res = max(res, beam_res)

    print(res)

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
