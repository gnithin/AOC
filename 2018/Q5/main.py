import re


def react(ip):
    curr_index = 0
    react_str = ip
    while curr_index < (len(react_str) - 1):
        curr_char = react_str[curr_index]
        next_char = react_str[curr_index + 1]
        if curr_char != next_char and curr_char.lower() == next_char.lower():
            end_val = curr_index + 2
            if end_val >= len(react_str):
                end_val = len(react_str) - 1
            react_str = react_str[:curr_index] + react_str[end_val:]
            if curr_index != 0:
                curr_index -= 1
        else:
            curr_index += 1
    return react_str


def get_react_len(ip):
    return len(react(ip))


def find_shortest_reaction(ip):
    original_ip = ip
    all_chars = set(original_ip.lower())

    min_len = 10000000
    for char in all_chars:
        new_ip = re.sub(char, '', original_ip, flags=re.IGNORECASE)
        curr_len = get_react_len(new_ip)
        if curr_len < min_len:
            min_len = curr_len
    return min_len


if __name__ == "__main__":
    filename = "ip.txt"
    # filename = "trial_ip.txt"

    with open(filename) as fp:
        contents = fp.read()

    ip = [ip.strip() for ip in contents.split("\n") if ip.strip() != ""][0]
    final_op = react(ip)
    print("After reaction - ", len(final_op))

    short_len = find_shortest_reaction(ip)
    print("Shortest len - ", short_len)

