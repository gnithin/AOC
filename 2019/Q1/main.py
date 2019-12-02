mem = {}


def get_fuel(ip):
    if ip in mem:
        return mem[ip]
    val = (ip // 3) - 2
    mem[ip] = val
    return val


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [int(line.strip()) for line in fp]

    # ip_list = [1969]
    mem = {}
    s = 0
    for ip in ip_list:
        fuel = get_fuel(ip)
        if fuel > 0:
            s += fuel
            while fuel > 0:
                fuel = get_fuel(fuel)
                if fuel > 0:
                    s += fuel
    print(s)
