from datetime import datetime, time
import re

TIME_FMT = "%Y-%m-%d %H:%M"


class Guard:
    def __init__(self, id):
        self.id = id
        self.shift_start_times = []
        self.shift_end_times = []
        self.sleep_start_times = []
        self.sleep_end_times = []

    def begins_shift_at(self, shift_time):
        self.shift_start_times.append(shift_time)

    def ends_shift_at(self, shift_time):
        self.shift_end_times.append(shift_time)

    def falls_asleep_at (self, sleep_time):
        self.sleep_start_times.append(sleep_time)

    def wakes_up_at(self, wake_time):
        self.sleep_end_times.append(wake_time)

    def get_total_sleep_time(self):
        if len(self.sleep_start_times) != len(self.sleep_end_times):
            print("Sleep start and end times don't match!")
            exit(1)

        total_sleep_time = 0
        for i in range(len(self.sleep_start_times)):
            start = self.sleep_start_times[i]
            end = self.sleep_end_times[i]
            total_sleep_time += int((end - start).seconds/60)
        return total_sleep_time

    def get_sleep_dist(self):
        minute_dict = {}
        for i in range(len(self.sleep_start_times)):
            start = self.sleep_start_times[i]
            end = self.sleep_end_times[i]

            start_min = start.minute
            end_min = end.minute

            count = start_min
            while count != end_min:
                if minute_dict.get(count, None) is None:
                    minute_dict[count] = 1
                else:
                    minute_dict[count] += 1
                count = (count + 1) % 60
        return minute_dict

    def get_max_sleep_min(self):
        minute_dict = self.get_sleep_dist()

        # Find the max entry here
        max_min = None
        max_entry = 0
        for k in minute_dict:
            if minute_dict[k] > max_entry:
                max_entry = minute_dict[k]
                max_min = k
        return max_min

    def get_min_most_slept_at(self):
        dist = self.get_sleep_dist()
        if len(dist) == 0:
            return 0, 0
        min = max(dist, key=dist.get)
        return min, dist[min]

    def __str__(self):
        return "\n".join([
            self.id,
            ",".join([s.strftime(TIME_FMT) for s in self.shift_start_times]),
            ",".join([s.strftime(TIME_FMT) for s in self.shift_end_times]),
            ",".join([s.strftime(TIME_FMT) for s in self.sleep_start_times]),
            ",".join([s.strftime(TIME_FMT) for s in self.sleep_end_times]),
        ])


class GuardManager:
    def __init__(self):
        self.guard_map = {}

    def get_guard_for(self, guard_id):
        guard = self.guard_map.get(guard_id)
        if guard is None:
            guard = Guard(guard_id)
            self.guard_map[guard_id] = guard
        return guard

    def __str__(self):
        g_list = self.guard_map.values()
        return "\n".join([g.__str__() for g in g_list])

    def get_guard_with_max_sleep(self):
        sel_guard_id = None
        sel_guard = None
        max_sleep = -1
        for guard_id in self.guard_map:
            guard = self.guard_map[guard_id]

            sleep_time = guard.get_total_sleep_time()
            # print("Guard id - ", guard_id, " sleep-time - ", sleep_time)
            if sleep_time > max_sleep:
                sel_guard_id = guard_id
                sel_guard = guard
                max_sleep = sleep_time
        return sel_guard_id, sel_guard

    def get_guard_sleep_index(self):
        guard_id, guard = self.get_guard_with_max_sleep()
        print("Fetching for guard - ", guard_id)
        max_sleep_min = guard.get_max_sleep_min()
        print("Max sleep - ", max_sleep_min)

        return int(guard_id.strip("#")) * max_sleep_min

    def get_least_productive_min_index(self):
        # Get the least prod minute
        sel_min = -1
        sel_guard_id = ""
        most_slept_at = -1

        for guard_id in self.guard_map:
            guard = self.guard_map[guard_id]
            min, freq = guard.get_min_most_slept_at()
            if freq > most_slept_at:
                most_slept_at = freq
                sel_min = min
                sel_guard_id = guard_id
        return int(sel_guard_id.strip("#")) * sel_min


if __name__ == "__main__":
    filename = "trial_ip.txt"
    filename = "ip.txt"

    with open(filename) as fp:
        contents = fp.read()

    ip_list = sorted([ip.strip() for ip in contents.split("\n") if ip.strip() != ""])
    manager = GuardManager()

    # Parse the ip
    regex = re.compile(r'^\[([^\]]+)\]\s*(?:(?:guard\s*(#[^ ]+))|(falls)|(wakes))', re.IGNORECASE)
    curr_guard_id = ""
    for ip in ip_list:
        matches = re.findall(regex, ip)
        if len(matches) == 0:
            print("Could not find any matches")
            exit(1)
        entries = matches[0]
        curr_time = datetime.strptime(entries[0], TIME_FMT)
        if entries[1] != "":
            if curr_guard_id != "":
                guard = manager.get_guard_for(curr_guard_id)
                guard.ends_shift_at(curr_time)

            # New id
            guard_id = entries[1]
            curr_guard_id = guard_id
            guard = manager.get_guard_for(curr_guard_id)
            guard.begins_shift_at(curr_time)
        elif entries[2] != "":
            # falls asleep
            guard = manager.get_guard_for(curr_guard_id)
            guard.falls_asleep_at(curr_time)
        elif entries[3] != "":
            # wakes up
            guard = manager.get_guard_for(curr_guard_id)
            guard.wakes_up_at(curr_time)
    index = manager.get_guard_sleep_index()
    print(index)

    index = manager.get_least_productive_min_index()
    print(index)
