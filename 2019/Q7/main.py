from itertools import permutations


class IntCode:
    def __init__(self, li, input_src):
        self.li = li
        self.input_src = input_src

    def process_test(self):
        i = 0
        while i < len(self.li):
            cmd = self.get_cmd(self.li[i])
            params = self.get_params(self.li[i])

            if cmd == 99:
                break
            elif cmd == 1:
                self.li[self.li[i + 3]] = self.get_val_for_mode(self.li, i + 1, params[0]) + self.get_val_for_mode(
                    self.li, i + 2,
                    params[1])
                i = i + 4
            elif cmd == 2:
                self.li[self.li[i + 3]] = self.get_val_for_mode(self.li, i + 1, params[0]) * self.get_val_for_mode(
                    self.li, i + 2,
                    params[1])
                i = i + 4
            elif cmd == 3:
                ip = int(next(self.input_src))
                self.li[self.li[i + 1]] = ip
                i += 2
            elif cmd == 4:
                op = self.li[self.li[i + 1]]
                yield op
                i += 2
            elif cmd == 5:
                ## Jump if true
                condn = self.get_val_for_mode(self.li, i + 1, params[0])
                if condn != 0:
                    loc = self.get_val_for_mode(self.li, i + 2, params[1])
                    i = loc
                else:
                    i += 3
            elif cmd == 6:
                ## Jump if false
                condn = self.get_val_for_mode(self.li, i + 1, params[0])
                if condn == 0:
                    loc = self.get_val_for_mode(self.li, i + 2, params[1])
                    i = loc
                else:
                    i += 3
            elif cmd == 7:
                ## less than
                p1 = self.get_val_for_mode(self.li, i + 1, params[0])
                p2 = self.get_val_for_mode(self.li, i + 2, params[1])
                val = 0
                if p1 < p2:
                    val = 1
                self.li[self.li[i + 3]] = val
                i = i + 4
            elif cmd == 8:
                ## Equal
                p1 = self.get_val_for_mode(self.li, i + 1, params[0])
                p2 = self.get_val_for_mode(self.li, i + 2, params[1])
                val = 0
                if p1 == p2:
                    val = 1
                self.li[self.li[i + 3]] = val
                i = i + 4

    def get_cmd(self, val):
        return val % 100

    def get_params(self, val):
        val = val // 100
        modes = []
        while val > 0:
            modes.append(val % 10)
            val = val // 10

        # Appending more values that might be necessary
        modes.append(0)
        modes.append(0)
        modes.append(0)
        modes.append(0)
        return modes

    def get_val_for_mode(self, li, index, mode):
        if mode == 0:
            return li[li[index]]
        else:
            return li[index]


def gen(li):
    for i in li:
        yield str(i)


def get_max_amplitude(ip_list, chain_size):
    outputs_map = {}
    for p in permutations(range(chain_size)):
        output_val = 0
        for phase in p:
            ip_func = gen([phase, output_val])
            ic_code = IntCode(ip_list[:], ip_func)
            output_val = next(ic_code.process_test())
        outputs_map[p] = output_val
    return outputs_map


if __name__ == "__main__":
    ip_list = []
    with open("ip1.txt", "r") as fp:
        for line in fp:
            for entry in line.strip().split(","):
                ip_list.append(int(entry))

    amplitudes = get_max_amplitude(ip_list, 5)
    print(max(amplitudes.values()))
