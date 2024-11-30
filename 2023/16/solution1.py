import fileinput

def solve(lines: list[str]) -> None:
    h, w = len(lines), len(lines[0])

    energy = [[0 for _ in range(w)] for _ in range(h)]

    # first beam, from 0,0 moving to right
    beams = [{
        "x": 0,
        "y": 0,
        "dx": 1,
        "dy": 0,
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
        
    res = 0 
    for e in energy:
        res += e.count(1)

    print(res)

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
