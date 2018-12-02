def find_checksum(ids):
    num_3 = 0
    num_2 = 0
    for entry in ids:
        print(entry)
        freq_map = find_freq(entry)
        print(freq_map)
        temp_freq_store = [0, 0]
        for c in freq_map:
            freq = freq_map[c]
            if freq > 2:
                temp_freq_store[1] = 1
            elif freq == 2:
                temp_freq_store[0] = 1
        num_2 += temp_freq_store[0]
        num_3 += temp_freq_store[1]
    return num_2 * num_3


def find_freq(entry):
    freq_map = {}
    for c in entry:
        if c in freq_map:
            freq_map[c] += 1
        else:
            freq_map[c] = 1
    return freq_map


if __name__ == "__main__":
    # file_name = "trial_ip.txt"
    file_name = "ip.txt"
    with open(file_name, 'r') as fp:
        contents = fp.read()
    ip_contents = [ip.strip() for ip in contents.strip().split("\n")
                   if ip.strip() != ""]
    print(ip_contents)
    checksum = find_checksum(ip_contents)
    print("Checksum -")
    print(checksum)
