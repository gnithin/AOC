from dataclasses import dataclass, field
from enum import Enum, auto
from typing import Generator, Dict

from Intcode import Gen, IntCode


class TileType(Enum):
    EMPTY = 0
    WALL = auto()
    BLOCK = auto()
    HOR_PADDLE = auto()
    BALL = auto()

    @staticmethod
    def get_type(i: int):
        l = [
            TileType.EMPTY,
            TileType.WALL,
            TileType.BLOCK,
            TileType.HOR_PADDLE,
            TileType.BALL,
        ]
        if i > len(l):
            raise Exception(f"Invalid type {i}!")
        return l[i]

    @staticmethod
    def get_char(type):
        l_map = {
            TileType.EMPTY: ".",
            TileType.WALL: "W",
            TileType.BLOCK: "B",
            TileType.HOR_PADDLE: "H",
            TileType.BALL: "O",
        }
        return l_map[type]


@dataclass()
class Coord:
    x: int = field(default=0)
    y: int = field(default=0)

    def __hash__(self):
        return 100 * self.x + self.y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __str__(self):
        return f"{self.x}, {self.y}"


@dataclass
class Arcade:
    ic_gen: Generator[int, None, None]
    generator: Gen
    __coords: Dict[Coord, TileType] = field(default_factory=dict)
    __score: int = field(default=0)

    def get_coords(self):
        while True:
            try:
                x = next(ic_gen)
                y = next(ic_gen)
                id = next(ic_gen)
                # print(f"Got - {x}, {y}, {id}")
                if x == -1 and y == 0:
                    self.__score = id
                else:
                    self.__coords[Coord(x, y)] = TileType.get_type(id)
                self.draw()

            except StopIteration:
                print(e)
            except Exception as e:
                print(e)
        return self.__coords

    def draw(self):
        max_x = 0
        max_y = 0
        min_x = 0
        min_y = 0
        for coords in self.__coords:
            if coords.x > max_x:
                max_x = coords.x
            if coords.y > max_y:
                max_y = coords.y
            if coords.x < min_x:
                min_x = coords.x
            if coords.y < min_y:
                min_y = coords.y

        # Increment for this to be used with range safely
        max_x += 1
        max_y += 1

        # X- cols
        # Y- rows
        matrix = [[0 for _ in range(max_x)] for _ in range(max_y)]

        for coords, type in self.__coords.items():
            matrix[coords.y][coords.x] = TileType.get_char(type)

        for i in range(max_y):
            print(" ".join(map(str, [matrix[i][j] for j in range(max_x)])))

        print(f"Score = {self.__score}")

    def play(self):
        self.get_coords()
        # self.draw()

        while True:
            # move_str = input("Your Move - (Left:-1, Right:1, Neutral:0) ")
            # move_val = int(move_str)
            # print("Loading move val - ", move_val)
            # self.generator.add_list_val(move_val)
            # print("*" * 10)

            self.get_coords()


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

    arcade = Arcade(ic_gen, generator)
    # Part 1
    # coords = arcade.get_coords()
    #
    # count = 0
    # for k, v in coords.items():
    #     if v == TileType.BLOCK:
    #         count += 1
    # print(count)
    arcade.play()
