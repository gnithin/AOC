import math
from typing import List, Dict


class Coord:
    def __init__(self, row, col):
        self.row = row
        self.col = col

    def __str__(self):
        return f"{self.col}, {self.row}"


class Line:
    def __init__(self, slope, y_intercept, dirn):
        self.slope = slope
        self.y_intercept = y_intercept
        self.dirn = dirn

    def get_tuple(self):
        return (self.slope, self.y_intercept, self.dirn)

    def __str__(self):
        return f"(m={self.slope}, c={self.y_intercept}, dirn={self.dirn})"


class SpaceMap:
    def __init__(self, ip):
        self.rows = 0
        self.cols = 0
        self.__asteroid_list: List[Coord] = []
        self.parse_input(ip)

    def parse_input(self, ip):
        self.rows = len(ip)
        self.cols = len(ip[0])
        for i in range(self.rows):
            for j in range(self.cols):
                if ip[i][j] == "#":
                    self.__asteroid_list.append(Coord(i, j))

    def __str__(self):
        s = "[\n"
        for e in self.__asteroid_list:
            s += f"\t{str(e)}\n"
        s += "]"
        return s

    def find_best_monitoring_station(self):
        visible_map = {}
        for asteroid_point in self.__asteroid_list:
            line_set = set()
            for other_point in self.__asteroid_list:
                if other_point == asteroid_point:
                    continue
                line_point = self.get_line_equation(asteroid_point, other_point)
                # line_set.add((line_point, other_point.__str__()))
                line_set.add(line_point.get_tuple())
            visible_map[asteroid_point] = line_set

        # Print the map
        # print("*" * 10)
        # for k in visible_map:
        #     print(f"{k} => {len(visible_map[k])}")
        #     # print(f"{k} => {(visible_map[k])}")
        # print("*" * 10)

        # Find the max entry
        max_val = 0
        max_key = None
        for k, v in visible_map.items():
            if len(v) > max_val:
                max_val = len(v)
                max_key = k
        return max_key, max_val

    @staticmethod
    def get_line_equation(point1: Coord, point2: Coord):
        deno_diff = point2.col - point1.col
        num_diff = point2.row - point1.row
        if deno_diff == 0:
            if num_diff > 0:
                m = math.inf
            else:
                m = -1 * math.inf
        else:
            m = (num_diff / deno_diff)
        c = point1.row - (m * point1.col)
        dirn = math.atan2(num_diff, deno_diff)
        return Line(m, c, dirn)

    def find_vaporized_order(self):
        pos, _ = self.find_best_monitoring_station()
        print(pos)

        # Find the equation between this line and every other line
        lines_map: Dict[Coord, Line] = {}
        for asteroid in self.__asteroid_list:
            if asteroid == pos:
                continue
            equation = self.get_line_equation(asteroid, pos)
            lines_map[asteroid] = equation

        dirn_map: Dict[float, List[Coord]] = {}
        for coord, line in lines_map.items():
            dirn = line.dirn
            if dirn not in dirn_map:
                dirn_map[dirn] = []
            dirn_map[dirn].append(coord)

        for k in dirn_map:
            dirn_map[k] = sorted(
                dirn_map[k],
                key=lambda x: self.find_distance(pos, x),
            )
        for k, v in dirn_map.items():
            print(f"k -> {k}, v -> {[str(coord) for coord in v]}")

        dirn_list = list(dirn_map.keys())
        formula = lambda x: x >= math.pi / 2
        rest = sorted(
            list(
                filter(
                    formula,
                    dirn_list
                )
            ),
        )
        second_quad = sorted(
            list(
                filter(
                    lambda x: not formula(x),
                    dirn_list
                )
            ),
        )
        rest.extend(second_quad)
        dirn_list = rest[:]

        # Go through each entry in the list and make sure to remove them one by one
        i = 0
        item_num = 0
        vaporized_list = []
        while len(dirn_map) > 0:
            dirn = dirn_list[i]
            if dirn in dirn_map:
                coord = dirn_map[dirn].pop(0)
                item_num += 1
                if len(dirn_map[dirn]) == 0:
                    del dirn_map[dirn]
                print(f"Remove item {item_num} -> {coord}")
                vaporized_list.append(coord)

            i = (i + 1) % len(dirn_list)
        return vaporized_list

    @staticmethod
    def find_distance(coord1: Coord, coord2: Coord) -> float:
        return ((coord1.col - coord2.col) ** 2 + (coord1.row - coord2.row) ** 2) ** 0.5


if __name__ == "__main__":
    # filename = "ip1.txt"
    filename = "ip2.txt"
    with open(filename, "r") as fp:
        ip_list = []
        for line in fp:
            new_list = []
            for c in line.strip():
                new_list.append(c)
            ip_list.append(new_list)

    # # Print the input
    # for row in ip_list:
    #     print(row)

    space_map = SpaceMap(ip_list)
    # print("Space contents - ")
    # print(space_map)
    #
    # # Part I
    # pos, num_visible_asteroids = space_map.find_best_monitoring_station()
    # print(f"Pos - {pos}, Num visible asteroids - {num_visible_asteroids}")

    # Part II
    vaporized_list = space_map.find_vaporized_order()
    e = vaporized_list[199]
    print(e.col * 100 + e.row)
