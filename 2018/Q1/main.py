def get_repeated_freq(seq):
    freq_list = [0]
    curr_freq = 0
    num_times = 0
    while True:
        print(num_times)
        num_times += 1
        for s in seq:
            curr_freq += s
            if curr_freq in freq_list:
                return curr_freq
            else:
                freq_list.append(curr_freq)


if __name__ == "__main__":
    with open("ip.txt", 'r') as fp:
        file_contents = fp.read()
    ip_entries = [int(s.strip(",")) for s in file_contents.strip().split("\n")]
    print(sum(ip_entries))

    # ip_entries = [+3, +3, +4, -2, -4]
    # ip_entries = [+7, +7, -2, -7, -4]
    print(get_repeated_freq(ip_entries))
