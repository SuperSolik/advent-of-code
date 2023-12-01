import fileinput
from typing import Any


def build_grammar(from_: list[tuple[str, int]]) -> dict[str, Any]:
    grammar = {}
    
    for (word, value) in from_:
        start = grammar
        for c in word[:-1]:
            start = start.setdefault(c, {})
        start[word[-1:]] = value

    return grammar


grammar_from = [
    ('one', 1), ('two', 2), ('three', 3), ('four', 4), ('five', 5), ('six', 6), ('seven', 7), ('eight', 8), ('nine', 9),
    ('0', 0), ('1', 1), ('2', 2), ('3', 3), ('4', 4), ('5', 5), ('6', 6), ('7', 7), ('8', 8), ('9', 9),
]


grammar = build_grammar(grammar_from)


def match(grammar, text) -> int | None:
    node = grammar
    for c in text:
        node = node.get(c)
        if node is None or isinstance(node, int):
            return node


def parse_line(line: str) -> int:
    pair = [None, None]

    for i in range(len(line)):
        r = match(grammar, line[i:])
        if r is not None:
            if pair[0] is None:
                pair[0] = r
            else:
                pair[1] = r

    if pair[1] is None:
        pair[1] = pair[0]

    return pair[0] * 10 + pair[1] if all(pair) else 0
    

if __name__ == '__main__':
    print(sum((parse_line(line) for line in fileinput.input())))
