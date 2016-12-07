import sys


def get_string(ip_list, part1=True):
    rotated_list = zip(*ip_list)
    final_string = ''
    for row in rotated_list:
        max_freq_dict = {}
        for char in row:
            max_freq_dict[char] = max_freq_dict.get(char, 0) + 1
        # print(max_freq_dict)

        if part1:
            max_freq = -1
            max_char = ''
            for char, freq in max_freq_dict.items():
                if freq > max_freq:
                    max_char = char
                    max_freq = freq
            final_string += max_char
        else:
            min_freq = 1000000
            min_char = ''
            for char, freq in max_freq_dict.items():
                if freq < min_freq:
                    min_char = char
                    min_freq = freq
            final_string += min_char
    return final_string

if __name__ == "__main__":
    ip_list = [ip.strip() for ip in sys.stdin]
    final_string = get_string(ip_list)
    print(final_string)

    modified_string = get_string(ip_list, False)
    print(modified_string)

