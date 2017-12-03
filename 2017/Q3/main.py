from __future__ import print_function


def build_matrix_of_size(side):
    if side % 2 == 0:
        return None

    matrix = list()
    for i in xrange(side):
        new_matrix = list()
        for j in xrange(side):
            new_matrix.append(0)
        matrix.append(new_matrix)

    x = side-1
    y = side-1
    dirn = "l"

    left_guard = 0
    right_guard = side
    top_guard = 0
    bottom_guard = side-1
    pos_list = list()

    for val in xrange(side**2):
        # print(str(x) + " - " + str(y))
        pos_list.append((y, x))
        new_x = x
        new_y = y
        if dirn == "l":
            new_x = x - 1
            if new_x < left_guard:
                new_x = x
                new_y = y - 1
                dirn = "u"
                left_guard += 1

        elif dirn == "r":
            new_x = x + 1
            if new_x == right_guard:
                new_x = x
                new_y = y + 1
                dirn = "d"
                right_guard -= 1

        elif dirn == "u":
            new_y = y - 1
            if new_y < top_guard:
                new_y = y
                new_x = x + 1
                dirn = "r"
                top_guard += 1

        else:
            new_y = y + 1
            if new_y == bottom_guard:
                new_y = y
                new_x = x - 1
                dirn = "l"
                bottom_guard -= 1
        x = new_x
        y = new_y

    val_counter = 1
    pos_list = pos_list[-1::-1]

    for x, y in pos_list:
        matrix[x][y] = val_counter
        val_counter += 1
    return matrix


def print_matrix(matrix):
    for i in xrange(len(matrix)):
        for j in xrange(len(matrix)):
            print('%3s' % matrix[i][j], end=" ")
        print("")


def get_nearest_distance_to_element(element):
    matrix_size = int((element ** 0.5) + 1)
    if matrix_size % 2 == 0:
        matrix_size += 1

    matrix = build_matrix_of_size(matrix_size)
    # print_matrix(matrix)
    pos_x = 0
    pos_y = 0
    for i in xrange(matrix_size):
        for j in xrange(matrix_size):
            if matrix[i][j] == element:
                pos_x = i
                pos_y = j

    print("x - " + str(pos_x) + " y - " + str(pos_y))
    centre_pos_x = centre_pos_y = (matrix_size - 1)/2
    num_steps = abs(pos_x - centre_pos_x) + abs(pos_y - centre_pos_y)

    return num_steps


if __name__ == "__main__":
    steps = get_nearest_distance_to_element(347991)
    print("steps - " + str(steps))
