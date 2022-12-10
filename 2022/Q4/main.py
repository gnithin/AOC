def does_comp_include_other(l, r):
    if l[0] < r[0]:
        return r[1] <= l[1]
    elif l[0] == r[0]:
        return True
    else:
        return l[1] <= r[1]


def does_overlap(l, r):
    if l[0] < r[0] and l[1] < r[0]:
        return False
    elif r[0] < l[0] and r[1] < l[0]:
        return False
    return True

if __name__ == "__main__":
    ip_list = []
    total = 0
    with open("ip2.txt", "r") as fp:
        for line in fp:
            ip = line.strip().split(",")
            comp = [e.split("-") for e in ip]
            ip_list.append(comp)
            if does_overlap(
                    list(map(lambda x: int(x), comp[0])),
                    list(map(lambda x: int(x), comp[1]))
            ):
                # print(line)
                total += 1

    print(total)
