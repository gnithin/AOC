if __name__ == "__main__":
    with open("ip.txt", 'r') as fp:
        file_contents = fp.read()
    ip_entries = [int(s.strip(",")) for s in file_contents.strip().split("\n")]
    print(sum(ip_entries))
