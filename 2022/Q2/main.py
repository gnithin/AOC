if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [line.strip().split(" ") for line in fp]
    print(ip_list)
