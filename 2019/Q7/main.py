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


class Gen:
    def __init__(self, li):
        self.li = li

    def gen(self):
        i = 0
        while i < len(self.li):
            yield self.li[i]
            i += 1

    def add_list_val(self, val):
        self.li.append(val)


def get_max_amplitude(ip_list, chain_size):
    outputs_map = {}
    for p in permutations(range(chain_size)):
        output_val = 0
        for phase in p:
            ip_func = Gen([phase, output_val]).gen()
            ic_code = IntCode(ip_list[:], ip_func)
            try:
                output_val = next(ic_code.process_test())
            except StopIteration:
                break
        outputs_map[p] = output_val
    return outputs_map


def get_max_amplitude_from_feedback(ip_list, chain_size):
    outputs_map = {}
    for p in permutations(range(chain_size, chain_size * 2)):
        output_val = 0
        did_stop = False

        test_codes = []
        ip_funcs = []
        for phase in p:
            ip_func = Gen([phase, output_val])
            ip_funcs.append(ip_func)

            ic_code = IntCode(ip_list[:], ip_func.gen())
            test_code = ic_code.process_test()
            test_codes.append(test_code)

            try:
                output_val = next(test_code)
            except StopIteration:
                did_stop = True
                break

        i = 0
        while not did_stop:
            ip_funcs[i].add_list_val(output_val)
            try:
                output_val = next(test_codes[i])
            except StopIteration:
                did_stop = True
                break
            i = (i + 1) % len(p)

        outputs_map[p] = output_val
    return outputs_map


if __name__ == "__main__":
    ip_list = []
    with open("ip2.txt", "r") as fp:
        for line in fp:
            for entry in line.strip().split(","):
                ip_list.append(int(entry))

    # amplitudes = get_max_amplitude(ip_list, 5)
    # print(max(amplitudes.values()))

    amplitudes = get_max_amplitude_from_feedback(ip_list, 5)
    print(max(amplitudes.values()))
