WIN = 6
DRAW = 3
LOSE = 0

# Rock, Paper, Scissor
LEFT = ["A", "B", "C"]
RIGHT = ["X", "Y", "Z"]

BEATS = [1, 2, 0]


def contest_points(l_index, r_index):
    if l_index == r_index:
        return DRAW
    elif BEATS[l_index] == r_index:
        return WIN
    return LOSE


def contest_points_v2(l_index, r_index):
    if r_index == 0:
        return LOSE + BEATS.index(l_index) + 1

    elif r_index == 1:
        return DRAW + (l_index + 1)

    return WIN + BEATS[l_index] + 1


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [line.strip().split(" ") for line in fp]

    points = 0
    for left, right in ip_list:
        t_points = contest_points_v2(LEFT.index(left), RIGHT.index(right))
        points += t_points
    print(points)
