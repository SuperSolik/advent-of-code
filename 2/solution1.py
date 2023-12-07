import fileinput


LIMITS = {
    "red": 12,
    "green": 13,
    "blue": 14
}

def check_game(line: str) -> int:
    game_data, rounds_data = line.split(':')
    game_number = int(game_data.split()[1])
    rounds_data = rounds_data.split(';')

    max_cnt_by_color = {}
    for round_record in rounds_data:
        for cube_count in round_record.strip().split(', '):
            cnt, color = cube_count.split()
            cnt = int(cnt)
                        
            # update max cnts
            cur_max_color_cnt = max_cnt_by_color.get(color, 0)
            max_cnt_by_color[color] = max(cur_max_color_cnt, int(cnt))

    
    for color, cnt in max_cnt_by_color.items():
        if cnt > LIMITS[color]:
            return 0

    return game_number

if __name__ == '__main__':
    print(sum(check_game(game) for game in fileinput.input()))
