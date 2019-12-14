import math


class Coord:
    def __init__(self, row, col):
        self.row = row
        self.col = col

    def __str__(self):
        return f"{self.col}, {self.row}"


class SpaceMap:
    def __init__(self, ip):
        self.rows = 0
        self.cols = 0
        self.__asteroid_list = []
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
                line_set.add(line_point)
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
        return c, dirn


if __name__ == "__main__":
    with open("ip2.txt", "r") as fp:
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
    pos, num_visible_asteroids = space_map.find_best_monitoring_station()
    print(f"Pos - {pos}, Num visible asteroids - {num_visible_asteroids}")
