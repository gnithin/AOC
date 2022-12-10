import re

if __name__ == "__main__":
    all_stacks = []

    def get_tops():
        s = ""
        for stack in all_stacks:
            s += stack[-1]
        return s

    def build_stacks(line):
        r = r'(\[.*?\]|\s{3})\s?'
        entries = re.findall(r, line)
        stack_index = 0
        for entry in entries:
            if len(all_stacks) < (stack_index + 1):
                all_stacks.append([])

            if entry.strip():
                curr_stack = all_stacks[stack_index]
                curr_stack.insert(0, entry.strip('[]'))
            stack_index += 1


    def move_stacks(pop_count, from_stack_index, to_stack_index):
        # print(pop_count, from_stack_index, to_stack_index)
        interim = []
        for _ in range(pop_count):
            interim.append(all_stacks[from_stack_index].pop())
        all_stacks[to_stack_index].extend(interim[::-1])


    with open("ip2.txt", "r") as fp:
        for line in fp:
            l = line.rstrip()
            if not l:
                continue
            elif l.lower().startswith("move"):
                r = r'move (\d+) from (\d+) to (\d+)'
                count, from_stack, to_stack = re.findall(r, l.strip())[0]
                move_stacks(int(count), int(from_stack) - 1, int(to_stack) - 1)

            elif '[' in l:
                build_stacks(l)
            else:
                continue

    # print(all_stacks)
    print(get_tops())
