import fileinput
import re

from operator import lt, gt


wf_pattern = re.compile(r'^(?P<name>\w+)\{(?P<rules>.*)\}$')
rule_pattern = re.compile(r'(?P<cond>(?P<label>[xmas])(?P<op>[><])(?P<value>\d+):)?(?P<dest>\w+)')
part_pattern = re.compile(r'^\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}$')

def parse_rule(rule: str) -> dict[str, str | None]:
    return rule_pattern.match(rule).groupdict()


def solve(lines: list[str]) -> None:
    wfs = {}
    parts = []
    wfs_done = False

    for line in lines:
        if line == "":
            wfs_done = True
            continue


        if not wfs_done:
            # it's a workflow
            parsed_wf = wf_pattern.match(line).groupdict()
            rules = []
            for rule in parsed_wf["rules"].split(','):
                parsed_rule = rule_pattern.match(rule).groupdict()
                if parsed_rule["cond"]:
                    parsed_rule["value"] = int(parsed_rule["value"])
                rules.append(parsed_rule)

            wfs[parsed_wf["name"]] = rules

        else:
            # it's a number

            x, m, a, s = part_pattern.match(line).groups()
            parts.append({
                "x": int(x),
                "m": int(m),
                "a": int(a),
                "s": int(s)
            })
    res = 0
    for p in parts:
        dest = "in"
        rules = wfs["in"][:]
        while dest not in "RA":
            rule = rules.pop(0)
            if not rule["cond"]:
                dest = rule["dest"]
                if dest not in "RA":
                    rules = wfs[dest][:]
            else:
                op = lt if rule["op"] == "<" else gt
                if op(p[rule["label"]], rule["value"]):
                    dest = rule["dest"]
                    if dest not in "RA":
                        rules = wfs[dest][:]
        
        if dest == "A":
            res += sum(p.values())

    print(res)

                
            


if __name__ == '__main__':
    solve([line.strip() for line in fileinput.input()])
