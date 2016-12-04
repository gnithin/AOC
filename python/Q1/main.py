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


def get_all_intermediate_coords(last_coord, ending_coord, prev_dirn):
    # Messy messy stuff
    # TODO: Clean this up!
    intermediate_values = []
    if prev_dirn in [EAST, WEST]:
        x2, x1 = ending_coord[0], last_coord[0]
        y = ending_coord[1]
        dirn_val = 1
        if prev_dirn == EAST:
            dirn_val = -1
        [
            intermediate_values.append((new_val, y))
            for new_val in range(x1 + dirn_val, x2 + dirn_val, dirn_val)
        ]
    else:
        y2, y1 = ending_coord[1], last_coord[1]
        x = ending_coord[0]
        dirn_val = 1
        if prev_dirn == SOUTH:
            dirn_val = -1
        [
            intermediate_values.append((x, new_val))
            for new_val in range(y1 + dirn_val, y2 + dirn_val, dirn_val)
        ]
    return intermediate_values


def get_shortest_path_len(coords, find_first_intersection=False):
    x_total = y_total = 0
    prev_dirn = NORTH
    all_visited_coords = [(0, 0)]

    for index, coord in enumerate(coords):
        # print("-"*20)
        # print(index, coord)

        mag, prev_dirn = get_next_mag_dirn(index, coord, prev_dirn)
        if index % 2:
            y_total += mag
        else:
            x_total += mag

        if find_first_intersection:
            new_coord = (x_total, y_total)
            prev_coord = all_visited_coords[-1]

            intermediate_values = get_all_intermediate_coords(
                prev_coord,
                new_coord,
                prev_dirn
            )

            intersection_flag = False
            for visited_coord in intermediate_values:
                if visited_coord in all_visited_coords:
                    intersection_flag = True
                    # print("For line between - ", prev_coord, new_coord)
                    # print(
                    # "Point in line for ", visited_coord,
                    # " is found? - ", intersection_flag
                    # )
                    break
            if intersection_flag:
                x_total, y_total = visited_coord
                break
            else:
                all_visited_coords.extend(intermediate_values)
            # print(all_visited_coords)

    return abs(x_total) + abs(y_total)

if __name__ == "__main__":
    ip_str = [ip.strip() for ip in sys.stdin][0]
    # ip_str = 'R5, L5, R5, R3'
    # ip_str = 'R8, R4, R4, R8'
    ip_list = ip_str.split(", ")

    shortest_len = get_shortest_path_len(ip_list)
    print("Shortest len - ", shortest_len)

    shortest_len_intersect = get_shortest_path_len(ip_list, True)
    print("Shortest len with intersections - ", shortest_len_intersect)

    print("For future changes")
    if (shortest_len_intersect, shortest_len) == (116, 241):
        print("All succeeded")
    else:
        print("All failed! :(")
