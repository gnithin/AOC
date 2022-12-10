def get_priority(c):
    o = ord(c.upper()) - ord('A') + 1
    if c.isupper():
        o += 26
    return o


def find_total(ip):
    first, second = ip[:len(ip) // 2], ip[len(ip) // 2:]
    common = list(set(first).intersection(set(second)))[0]
    return get_priority(common)


def find_group(ip_list):
    return get_priority(list(set.intersection(*[set(ip) for ip in ip_list]))[0])


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [line.strip() for line in fp]
    total = 0
    for i in range(0, len(ip_list), 3):
        total += find_group(ip_list[i:i + 3])
    print(total)
