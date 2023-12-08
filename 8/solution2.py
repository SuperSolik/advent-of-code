import fileinput
import re
import math

from itertools import cycle
from functools import reduce


node_pattern = re.compile(r'(\w+)\s=\s\((\w+),\s?(\w+)\)')


def solve(lines: list[str]) -> None:
    instructions, lines = lines[0].strip(), lines[2:]
    graph = {}

    start_nodes = []


    for line in lines:
        node, left, right = node_pattern.match(line).groups()
        if node.endswith('A'):
            start_nodes.append(node)

        graph[node] = {
            'R': right,
            'L': left
        }

    nodes_steps = []

    # calc steps for each node
    for node in start_nodes:        
        steps = 0
        actions = cycle(instructions)

        while not node.endswith('Z'):
            ins = next(actions)
            node = graph[node][ins]
            steps += 1

        nodes_steps.append(steps)

    # LCM
    print(reduce(lambda x, y: x * y //math.gcd(x, y), nodes_steps))
    

if __name__ == '__main__':
    solve([line for line in fileinput.input()])
