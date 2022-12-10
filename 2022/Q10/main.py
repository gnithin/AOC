class Console:
    NOOP = "noop"
    ADDX = "addx"

    def __init__(self, commands):
        self.register_x = 1
        self.tick = 1
        self.commands = commands
        self.signal_strength = []
        self.signal_strength_ticks = [20, 60, 100, 140, 180, 220]
        self.crt = ""

    def run(self):
        curr_cmd_index = 0
        wait_tick = 0
        prev_cmd = None
        while True:
            # Record signal strength
            if self.tick in self.signal_strength_ticks:
                self.signal_strength.append(self.tick * self.register_x)

            print(f"{self.tick}\t{self.register_x}\t{self.commands[curr_cmd_index]}")

            if wait_tick != 0:
                wait_tick -= 1
            else:
                command = self.commands[curr_cmd_index]
                if command[0] == self.NOOP:
                    curr_cmd_index += 1
                elif command[0] == self.ADDX:
                    wait_tick = 1
                prev_cmd = command

            if ((self.tick - 1) % 40) in [self.register_x - 1, self.register_x, self.register_x + 1]:
                self.add_to_crt(is_lit=True)
            else:
                self.add_to_crt(is_lit=False)

            self.tick += 1

            # Update logic
            if prev_cmd[0] == self.ADDX and wait_tick == 0:
                self.register_x += prev_cmd[1]
                curr_cmd_index += 1

            if curr_cmd_index >= len(self.commands):
                break

    def get_total_signal_strength(self) -> int:
        print(self.signal_strength)
        return sum(self.signal_strength)

    def add_to_crt(self, is_lit=True):
        self.crt += "#" if is_lit else "."

    def draw(self):
        # 6x40
        for i in range(6):
            start = i * 40
            print(self.crt[start:start + 40])


if __name__ == "__main__":
    commands = []
    with open("ip2.txt", "r") as fp:
        for line in fp:
            cmd = line.strip().split()
            if len(cmd) > 1:
                cmd[1] = int(cmd[1])
            commands.append(cmd)
        print(commands)

    console = Console(commands)
    console.run()
    print(console.get_total_signal_strength())
    console.draw()
