def create_points(ip):
    points_set = []
    coord = (0, 0)
    for (dirn, mag) in ip:
        if dirn == "U":
            [points_set.append((coord[0], coord[1] + i)) for i in range(1, mag + 1)]
            coord = (coord[0], coord[1] + mag)
        elif dirn == "D":
            [points_set.append((coord[0], coord[1] - i)) for i in range(1, mag + 1)]
            coord = (coord[0], coord[1] - mag)
        elif dirn == "R":
            [points_set.append((coord[0] + i, coord[1])) for i in range(1, mag + 1)]
            coord = (coord[0] + mag, coord[1])
        elif dirn == "L":
            [points_set.append((coord[0] - i, coord[1])) for i in range(1, mag + 1)]
            coord = (coord[0] - mag, coord[1])
    return points_set


def get_small_intersection(p1, p2):
    common = set(p1).intersection(set(p2))
    print(common)
    small = None
    min_val = None
    for p in common:
        s = abs(p[0]) + abs(p[1])
        if min_val is None or s < min_val:
            min_val = s
            small = p
    return small, min_val


def get_min_intersection(p1, p2):
    common = set(p1).intersection(set(p2))
    small = None
    min_val = None
    for p in common:
        s = p1.index(p) + p2.index(p) + 2
        if min_val is None or s < min_val:
            min_val = s
            small = s
    return small, min_val


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [line.strip().split(",") for line in fp]
        formatter = (lambda x: [(ip[0], int(ip[1:])) for ip in x])
        p1, p2 = formatter(ip_list[0]), formatter(ip_list[1])
    print(p1)
    print(p2)

    p1_points = create_points(p1)
    p2_points = create_points(p2)
    print(p1_points)
    print(p2_points)

    # print(get_small_intersection(p1_points, p2_points))
    print(get_min_intersection(p1_points, p2_points))
