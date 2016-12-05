import sys

DEFAULT_NUM_PAD = ['123', '456', '789']
INVALID_CHAR = '*'


def get_final_code(ip_str):
    curr_pos = (1, 1)
    final_code = ''
    for ip in ip_str:
        curr_code, curr_pos = get_numpad_code(ip, curr_pos)
        final_code += curr_code
    return final_code


def get_numpad_code(ip, initial_pos=(1, 1), custom_num_pad=None):
    print(ip)
    move_map = {
        'U': (0, -1),
        'D': (0, 1),
        'L': (-1, 0),
        'R': (1, 0),
    }

    if custom_num_pad is None:
        num_pad = DEFAULT_NUM_PAD

    max_num_pad_len = len(num_pad[0])
    curr_pos = initial_pos
    for dirn in ip:
        delta = move_map[dirn]
        new_pos = map(sum, zip(curr_pos, delta))

        invalid = False
        for pos in new_pos:
            if pos < 0 or pos > max_num_pad_len - 1:
                invalid = True
                break

        if (
            not invalid and
            num_pad[curr_pos[1]][curr_pos[0]] != INVALID_CHAR
        ):
                curr_pos = new_pos
                print(curr_pos)
        else:
            print("Invalid")

    curr_code = num_pad[curr_pos[1]][curr_pos[0]]
    print("Finally - ", (curr_code, curr_pos))

    return (curr_code, curr_pos)


if __name__ == "__main__":
    ip_str = [ip.strip() for ip in sys.stdin]
    final_code = get_final_code(ip_str)
    print("Final code -", final_code)
