SIZE = 14
def get_first_pos(code):
    for i in range(len(code) - SIZE):
        # print(code[i:i+4])
        if len(set(code[i:i+SIZE])) == SIZE:
            return i + SIZE
    return -1

if __name__ == "__main__":
    with open("ip2.txt", "r") as fp:
        ip_list = [line.strip() for line in fp]
        ip = ip_list[0]
        print(get_first_pos(ip))


