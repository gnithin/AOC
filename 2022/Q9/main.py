class RopeController:
    def __init__(self):
        # (row, col)
        self.head_pos = [0, 0]
        self.tail_pos = [0, 0]
        self.all_tail_positions = [(self.tail_pos[0], self.tail_pos[1])]

    def set_head_position(self, row, col):
        self.head_pos = [row, col]
        self._update_tail_position()

    def move_head(self, direction: str, amount: int):
        for step in range(amount):
            if direction == "U":
                self.head_pos[0] += 1
            elif direction == "D":
                self.head_pos[0] -= 1
            elif direction == "R":
                self.head_pos[1] += 1
            elif direction == "L":
                self.head_pos[1] -= 1
            self._update_tail_position()

    def _get_distance_tail_and_head(self):
        return [
            # if +ve, then head is one the right, else down
            self.head_pos[0] - self.tail_pos[0],
            # if +ve, then head is up, else down
            self.head_pos[1] - self.tail_pos[1]
        ]

    def _is_tail_touching_head(self) -> bool:
        r, c = self._get_distance_tail_and_head()
        return abs(r) <= 1 and abs(c) <= 1

    def _update_tail_position(self):
        if self._is_tail_touching_head():
            # Don't have to do anything
            return

        r_delta, c_delta = self._get_distance_tail_and_head()
        if abs(r_delta) > 0 and abs(c_delta) > 0:
            # MUST move diagonally (shortest manhattan distance doesn't count)
            self.tail_pos[0] += r_delta // abs(r_delta)
            self.tail_pos[1] += c_delta // abs(c_delta)
        elif abs(r_delta) > 1:
            self.tail_pos[0] += r_delta // abs(r_delta)
        elif abs(c_delta) > 1:
            self.tail_pos[1] += c_delta // abs(c_delta)
        self.all_tail_positions.append((self.tail_pos[0], self.tail_pos[1]))

    def get_total_tail_unique_visits(self) -> int:
        print(self.all_tail_positions)
        return len(set(self.all_tail_positions))


if __name__ == "__main__":
    ropes = []
    for i in range(10):
        ropes.append(RopeController())

    with open("ip1.txt", "r") as fp:
        for line in fp:
            direction, amount_str = line.strip().split(" ")
            amount = int(amount_str)
            for _ in range(amount):
                ropes[0].move_head(direction, 1)
                for i in range(1, 10):
                    if i == 1:
                        pos = ropes[0].head_pos
                    else:
                        pos = ropes[i - 1].tail_pos
                    ropes[i].set_head_position(*pos)

    print(ropes[9].get_total_tail_unique_visits())
