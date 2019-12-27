from typing import Generator, List


class CustomList:
    def __init__(self, li):
        self.map = {}
        self.__update_map_for_list(li)

    def __update_map_for_list(self, li):
        for i, e in enumerate(li):
            self.__update_map(i, e)

    def __update_map(self, index, entry):
        self.map[index] = entry

    def __getitem__(self, key):
        if key < 0:
            raise Exception("Negative index!")

        if key in self.map:
            return self.map[key]

        # Think about this
        return 0

    def __setitem__(self, key, value):
        if key < 0:
            # print("Negative index!")
            return
        self.map[key] = value

    def __len__(self):
        return len(self.map)


class IntCode:
    def __init__(self, li, input_src):
        self.li = CustomList(li)
        self.input_src = input_src
        self.relative_base = 0

    def process_test(self) -> Generator[int, None, None]:
        i = 0
        while i < len(self.li):
            cmd = self.get_cmd(self.li[i])
            params = self.get_params(self.li[i])
            # print("Command - ", cmd)
            # print("List - ", self.li.map)
            # print("Relative base - ", self.relative_base)
            # print("*" * 10)

            if cmd == 99:
                break
            elif cmd == 1:
                self.li[self.get_index_for_mode(self.li, i + 3, params[2])] = self.get_val_for_mode(self.li, i + 1,
                                                                                                    params[0]) + \
                                                                              self.get_val_for_mode(self.li, i + 2,
                                                                                                    params[1])
                i = i + 4
            elif cmd == 2:
                self.li[self.get_index_for_mode(self.li, i + 3, params[2])] = self.get_val_for_mode(self.li, i + 1,
                                                                                                    params[0]) * \
                                                                              self.get_val_for_mode(self.li, i + 2,
                                                                                                    params[1])
                i = i + 4
            elif cmd == 3:
                # ip = int(input("Ip - "))
                ip = int(next(self.input_src))
                # print("Val - " + str(ip))
                self.li[self.get_index_for_mode(self.li, i + 1, params[0])] = ip
                i += 2
            elif cmd == 4:
                op = self.get_val_for_mode(self.li, i + 1, params[0])
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
                self.li[self.get_index_for_mode(self.li, i + 3, params[2])] = val
                i = i + 4
            elif cmd == 8:
                ## Equal
                p1 = self.get_val_for_mode(self.li, i + 1, params[0])
                p2 = self.get_val_for_mode(self.li, i + 2, params[1])
                val = 0
                if p1 == p2:
                    val = 1
                self.li[self.get_index_for_mode(self.li, i + 3, params[2])] = val
                i = i + 4
            elif cmd == 9:
                # Relative base offset
                p1 = self.get_val_for_mode(self.li, i + 1, params[0])
                self.relative_base += p1
                i = i + 2

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
        modes.append(0)
        return modes

    def get_val_for_mode(self, li, index, mode):
        return li[self.get_index_for_mode(li, index, mode)]

    def get_index_for_mode(self, li, index, mode):
        if mode == 0:
            return li[index]
        elif mode == 2:
            return self.relative_base + li[index]
        else:
            return index


class Gen:
    def __init__(self, li):
        self.li: List[int] = li

    def gen(self) -> Generator[int, None, None]:
        i = 0
        while i < len(self.li):
            yield self.li[i]
            i += 1

    def add_list_val(self, val):
        # print("Added value - " + str(val))
        self.li.append(val)
