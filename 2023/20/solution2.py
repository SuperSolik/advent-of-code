import fileinput
import math

from functools import reduce
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
    
    rx_from = [name for name, module in modules.items() if "rx" in module.dest][0]
    rx_from_from = {name: {"found": False, "presses": 0} for name in modules[rx_from].state.keys()}

    i = 0
    while True:
        q = deque([])
        for d in modules["broadcaster"].dest:
            q.append(("broadcaster", d, "low"))

        while q:
            from_, to_, pulse = q.popleft()

            if to_ not in modules:
                continue

            if from_ in rx_from_from and pulse == "high":
                if not rx_from_from[from_]["found"]:
                    rx_from_from[from_]["presses"] = i
                    rx_from_from[from_]["found"] = True
                    break

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

        if all(m["found"] for m in rx_from_from.values()):
            break


        i += 1


    print(reduce(math.lcm, (m["presses"] + 1 for m in rx_from_from.values()), 1))


if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
