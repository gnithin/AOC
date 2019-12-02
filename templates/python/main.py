if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [int(line.strip()) for line in fp]
    print(ip_list)
