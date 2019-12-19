from dataclasses import dataclass, field
from enum import Enum, auto
from typing import Generator, Dict

from Intcode import IntCode, Gen


@dataclass
class Coord:
    x: int
    y: int

    def __hash__(self):
        return self.x * 100 + self.y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y


class Direction(Enum):
    UP = 0
    DOWN = auto()
    RIGHT = auto()
    LEFT = auto()


@dataclass
class HullRobot:
    ic_gen: Generator[int, None, None]
    gen: Gen
    __curr_pos: Coord = field(default=Coord(0, 0))
    __color_coords: Dict[Coord, int] = field(default_factory=dict)
    __curr_dirn: Direction = field(default=Direction.UP)
    initial = True

    def __get_color(self) -> int:
        if self.__curr_pos in self.__color_coords:
            return self.__color_coords[self.__curr_pos]
        if self.initial:
            self.initial = False
            return 1
        return 0

    def find_min_paint(self):
        while True:
            try:
                self.gen.add_list_val(self.__get_color())
                color_val = next(self.ic_gen)
                self.__color_coords[self.__curr_pos] = color_val

                dirn = next(self.ic_gen)
                self.__update_new_coord(dirn, 1)

            except StopIteration:
                break
        # print(self.__color_coords)
        return len(self.__color_coords)

    def __update_new_coord(self, dirn: int, step: int):
        if dirn == 0:
            # Turn left 90 degrees
            if self.__curr_dirn == Direction.UP:
                self.__curr_dirn = Direction.LEFT
            elif self.__curr_dirn == Direction.DOWN:
                self.__curr_dirn = Direction.RIGHT
            elif self.__curr_dirn == Direction.RIGHT:
                self.__curr_dirn = Direction.UP
            else:
                self.__curr_dirn = Direction.DOWN

        elif dirn == 1:
            # Turn right 90 degrees
            if self.__curr_dirn == Direction.UP:
                self.__curr_dirn = Direction.RIGHT
            elif self.__curr_dirn == Direction.DOWN:
                self.__curr_dirn = Direction.LEFT
            elif self.__curr_dirn == Direction.RIGHT:
                self.__curr_dirn = Direction.DOWN
            else:
                self.__curr_dirn = Direction.UP

        # Change the position
        if self.__curr_dirn == Direction.UP:
            self.__curr_pos = Coord(
                self.__curr_pos.x,
                self.__curr_pos.y + step
            )
        elif self.__curr_dirn == Direction.DOWN:
            self.__curr_pos = Coord(
                self.__curr_pos.x,
                self.__curr_pos.y - step
            )
        elif self.__curr_dirn == Direction.LEFT:
            self.__curr_pos = Coord(
                self.__curr_pos.x - step,
                self.__curr_pos.y
            )
        else:
            self.__curr_pos = Coord(
                self.__curr_pos.x + step,
                self.__curr_pos.y
            )

    def paint(self):
        # Normalize the coords
        coords = self.__color_coords.keys()
        min_x = min(coords, key=lambda c: c.x).x
        min_y = min(coords, key=lambda c: c.y).y
        print(min_x)
        print(min_y)
        normalized_coords = {}
        for coords, color in self.__color_coords.items():
            normalized_coords[Coord(coords.x + (-1) * min_x, coords.y + -1 * min_y)] = color

        max_x = (max(normalized_coords, key=lambda c: c.x).x) + 1
        max_y = (max(normalized_coords, key=lambda c: c.y).y) + 1

        matrix = []
        for x in range(max_x):
            matrix.append([0 for y in range(max_y)])

        for coord in normalized_coords:
            matrix[coord.x][coord.y] = normalized_coords[coord]

        # transpose the matrix
        # print the matrix
        for i in range(max_x):
            print("".join(map(str, [matrix[i][j] if matrix[i][j] == 1 else " " for j in range(max_y)])))


if __name__ == "__main__":
    filename = "ip1.txt"
    # filename = "ip2.txt"
    with open(filename, "r") as fp:
        lines = [line.strip() for line in fp]
        line = lines[0]
        ip_list = list(map(int, line.split(",")))

    generator = Gen([])
    ic_code = IntCode(ip_list, generator.gen())
    ic_gen = ic_code.process_test()

    robot = HullRobot(ic_gen, generator)
    min_paint = robot.find_min_paint()
    # print(min_paint)

    # Part 2
    robot.paint()
