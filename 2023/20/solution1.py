import fileinput
from collections import deque


class Module:
    def __init__(self, name: str, t: str, dest: list[str]):
        self.name = name
        self.t = t
        self.dest = dest
        self.state = None

        match t:
            case '%':
                self.state = False
            case '&':
                self.state = {}

    def __str__(self) -> str:
        return f"{self.name}({self.state}) -> {self.dest}"


def solve(lines: list[str]) -> None:
    modules = {}

    for line in lines:
        name, destinations = line.split(' -> ')
        if name.startswith(('%', '&')):
            t, name = name[0], name[1:]
        else:
            t, name = None, name
        modules[name] = Module(name, t, destinations.split(', '))

    for name, module in modules.items():
        for d in module.dest:
            if d in modules and modules[d].t == "&":
                modules[d].state[name] = "low"
    
    
    low_cnt, high_cnt = 0, 0

    for _ in range(1000):
        low_cnt += 1

        q = deque([])
        for d in modules["broadcaster"].dest:
            q.append(("broadcaster", d, "low"))

        while q:
            from_, to_, pulse = q.popleft()
            
            if pulse == "low":
                low_cnt += 1
            else:
                high_cnt += 1

            if to_ not in modules:
                continue

            new_pulse = None

            match modules[to_].t:
                case '&':
                    modules[to_].state[from_] = pulse
                    if all(p == "high" for p in modules[to_].state.values()):
                        new_pulse = "low"
                    else:
                        new_pulse = "high"

                case '%':
                    if pulse == "low":
                        modules[to_].state = not modules[to_].state
                        new_pulse = "high" if modules[to_].state else "low"
                    else:
                        continue

            for d in modules[to_].dest:
                q.append((to_, d, new_pulse))

    print(high_cnt, low_cnt)
    print(high_cnt * low_cnt)

if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
