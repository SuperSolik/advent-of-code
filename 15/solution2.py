import fileinput
import re


def xmas_hash(s: str) -> int:
    h = 0
    for ch in s:
        h += ord(ch)
        h *= 17
        h = h % 256
    return h


def solve(lines: list[str]) -> None:
    boxes = [[] for _ in range(256)]

    for p in lines[0].split(','):
        if "=" in p:
            # add
            label, new_f_len = p.split('=')
            h = xmas_hash(label) 
                
            c = 0
            for i, (cur_label, old_f_len) in enumerate(boxes[h]):
                if label == cur_label:
                    # existing entry
                    c += 1
                    boxes[h][i] = (label, int(new_f_len))
            if c <= 0:
                # no existing entries
                boxes[h].append((label, int(new_f_len)))
        else:
            # remove
            label = p[:-1] # the last is -
            h = xmas_hash(label) 

            # drop the label from the box
            boxes[h] = [(cur_label, f_len) for (cur_label, f_len) in boxes[h] if cur_label != label]

    res = 0
    for i, box in enumerate(boxes):
        res += sum((i + 1) * (j + 1) * f_len for j, (_, f_len) in enumerate(box))

    print(res)


if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
