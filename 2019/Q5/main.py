def process_test(li):
    i = 0
    while i < len(li):
        cmd = get_cmd(li[i])
        params = get_params(li[i])

        if cmd == 99:
            return li
        elif cmd == 1:
            li[li[i + 3]] = get_val_for_mode(li, i + 1, params[0]) + get_val_for_mode(li, i + 2, params[1])
            i = i + 4
        elif cmd == 2:
            li[li[i + 3]] = get_val_for_mode(li, i + 1, params[0]) * get_val_for_mode(li, i + 2, params[1])
            i = i + 4
        elif cmd == 3:
            ip = int(input("Input - "))
            li[li[i + 1]] = ip
            i += 2
        elif cmd == 4:
            op = li[li[i + 1]]
            print("output - ", op)
            i += 2
        elif cmd == 5:
            ## Jump if true
            condn = get_val_for_mode(li, i + 1, params[0])
            if condn != 0:
                loc = get_val_for_mode(li, i + 2, params[1])
                i = loc
            else:
                i += 3
        elif cmd == 6:
            ## Jump if false
            condn = get_val_for_mode(li, i + 1, params[0])
            if condn == 0:
                loc = get_val_for_mode(li, i + 2, params[1])
                i = loc
            else:
                i += 3
        elif cmd == 7:
            ## less than
            p1 = get_val_for_mode(li, i + 1, params[0])
            p2 = get_val_for_mode(li, i + 2, params[1])
            val = 0
            if p1 < p2:
                val = 1
            li[li[i + 3]] = val
            i = i + 4
        elif cmd == 8:
            ## Equal
            p1 = get_val_for_mode(li, i + 1, params[0])
            p2 = get_val_for_mode(li, i + 2, params[1])
            val = 0
            if p1 == p2:
                val = 1
            li[li[i + 3]] = val
            i = i + 4

    return li


def get_cmd(val):
    return val % 100


def get_params(val):
    val = val // 100
    modes = []
    while val > 0:
        modes.append(val % 10)
        val = val // 10

    # Appending more values that might be necessary
    modes.append(0)
    modes.append(0)
    modes.append(0)
    modes.append(0)
    return modes


def get_val_for_mode(li, index, mode):
    if mode == 0:
        return li[li[index]]
    else:
        return li[index]


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [int(i) for line in fp for i in line.split(",")]
    print(process_test(ip_list))
