import sys


def get_shortest_path_len(coords):
    NORTH, SOUTH, EAST, WEST = 0, 1, 2, 3

    x_total = 0
    y_total = 0
    curr_dirn = NORTH

    for index, coord in enumerate(coords):
        (dirn, mag) = coord[0], int(coord[1:])

        RCOMPARE, LCOMPARE, YCOMPARE, NCOMPARE = (
            EAST, WEST, SOUTH, NORTH
        ) if index % 2 else (
            NORTH, SOUTH, EAST, WEST
        )

        if (dirn, curr_dirn) in [('R', RCOMPARE), ('L', LCOMPARE)]:
            mag = -1 * mag
            curr_dirn = YCOMPARE
        else:
            curr_dirn = NCOMPARE

        # print("Mag - " + str(mag) + " dirn " + str(curr_dirn))

        if index % 2:
            # Evaluating vertical movement
            y_total += mag
        else:
            # Evaluating horizontal movement
            x_total += mag

    return abs(x_total) + abs(y_total)

if __name__ == "__main__":
    ip_str = [ip.strip() for ip in sys.stdin][0]
    # ip_str = 'R5, L5, R5, R3'
    ip_list = ip_str.split(", ")
    # print(ip_list)
    shortest_len = get_shortest_path_len(ip_list)
    print("Shortest len - ", shortest_len)
