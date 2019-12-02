def process(li):
    i = 0
    while i < len(li):
        cmd = li[i]
        if cmd == 99:
            return li
        elif cmd == 1:
            li[li[i + 3]] = li[li[i + 1]] + li[li[i + 2]]
        else:
            li[li[i + 3]] = li[li[i + 1]] * li[li[i + 2]]
        i = i + 4
    return li


def find_noun_verb(li, val):
    for noun in range(1, 100):
        for verb in range(1, 100):
            new_li = li[:]
            new_li[1] = noun
            new_li[2] = verb
            res = process(new_li)
            if res[0] == val:
                return (100 * noun) + verb
    return None


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        for line in fp:
            ip_list = [int(i) for i in line.split(",")]
    # print(process(ip_list))
    print(find_noun_verb(ip_list, 19690720))
