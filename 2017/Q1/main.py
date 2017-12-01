import sys


def find_sum_of_repeat_digits_in_seq_str(seq_str, consecutive_distance):
    seq_str_len = len(seq_str)
    computed_sum = 0
    for curr_index in xrange(seq_str_len):
        next_index = (curr_index+consecutive_distance) % seq_str_len
        if seq_str[curr_index] == seq_str[next_index]:
            computed_sum += int(seq_str[curr_index])
    return computed_sum


if __name__ == "__main__":
    ip = [ip.strip() for ip in sys.stdin if ip.strip() != ""][0]
    # P1
    sum_val = find_sum_of_repeat_digits_in_seq_str(ip, 1)
    print("P1: The sum of repeated digits - " + str(sum_val))
    # P2
    sum_val = find_sum_of_repeat_digits_in_seq_str(ip, len(ip)/2)
    print("P2: The sum of repeated digits - " + str(sum_val))
