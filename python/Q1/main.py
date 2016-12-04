import sys
NORTH, SOUTH, EAST, WEST = 0, 1, 2, 3


def get_next_mag_dirn(index, coord, prev_dirn):
    (dirn, mag) = coord[0], int(coord[1:])

    RCOMPARE, LCOMPARE, YCOMPARE, NCOMPARE = (
        EAST, WEST, SOUTH, NORTH
    ) if index % 2 else (
        NORTH, SOUTH, EAST, WEST
    )

    if (dirn, prev_dirn) in [('R', RCOMPARE), ('L', LCOMPARE)]:
        mag = -1 * mag
        prev_dirn = YCOMPARE
    else:
        prev_dirn = NCOMPARE
    return (mag, prev_dirn)


def get_shortest_path_len(coords):
    x_total = y_total = 0
    prev_dirn = NORTH

    for index, coord in enumerate(coords):
        mag, prev_dirn = get_next_mag_dirn(index, coord, prev_dirn)
        if index % 2:
            y_total += mag
        else:
            x_total += mag

    return abs(x_total) + abs(y_total)

if __name__ == "__main__":
    ip_str = [ip.strip() for ip in sys.stdin][0]
    # ip_str = 'R5, L5, R5, R3'
    # ip_str = 'R8, R4, R4, R8'
    ip_list = ip_str.split(", ")
    # print(ip_list)
    shortest_len = get_shortest_path_len(ip_list)
    print("Shortest len - ", shortest_len)
