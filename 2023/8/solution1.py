import fileinput
import re

from itertools import cycle


node_pattern = re.compile(r'(\w+)\s=\s\((\w+),\s?(\w+)\)')


def solve(lines: list[str]) -> None:
    instructions, lines = lines[0].strip(), lines[2:]
    graph = {}

    for line in lines:
        node, left, right = node_pattern.match(line).groups()
        graph[node] = {
            'R': right,
            'L': left
        }

    start = 'AAA'
    end = 'ZZZ'
    node = start

    instructions = cycle(instructions)

    steps = 0
    while node != end:
        i = next(instructions)
        node = graph[node][i]
        steps += 1

    print(steps)

    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
