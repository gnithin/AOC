import re


class Rect:
    def __init__(self, name, x, y, w, h):
        self.x = x
        self.y = y
        self.w = w
        self.h = h
        self.name = name

    def __str__(self):
        return ", ".join([
            self.name,
            str(self.x),
            str(self.y),
            str(self.w),
            str(self.h),
        ])

    def get_all_points(self):
        return [(x, y) for x in range(self.x, self.x + self.w)
                for y in range(self.y, self.y + self.h)]

    @staticmethod
    def get_rect_from_entry(entry):
        # #1 @ 1,3: 4x4
        all_matches = re.findall(r'(#\d+)\s+@\s+(\d+),(\d+):\s+(\d+)x(\d+)',
                                 entry)
        matches = all_matches[0]
        return Rect(
            matches[0],
            int(matches[1]),
            int(matches[2]),
            int(matches[3]),
            int(matches[4]),
        )


class RectManager:
    def __init__(self, rect_list):
        self.rect_list = rect_list

    def __str__(self):
        return "\n".join([rect.__str__() for rect in self.rect_list])

    def find_overlaps(self):
        mul_points_list = [rect.get_all_points() for rect in self.rect_list]
        max_entries = len(mul_points_list)
        overlaps = set()
        for i in range(0, max_entries - 1):
            for j in range(i+1, max_entries):
                for entry in \
                        set(
                            mul_points_list[i]).intersection(
                            set(mul_points_list[j])
                        ):
                    overlaps.add(entry)
        return overlaps

    def find_no_overlap_id(self):
        mul_points_list = [rect.get_all_points() for rect in self.rect_list]
        ids_list = [rect.name for rect in self.rect_list]
        max_entries = len(self.rect_list)
        for i in range(0, max_entries - 1):
            for j in range(i+1, max_entries):
                if self.rect_list[i].name not in ids_list and \
                        self.rect_list[j].name not in ids_list:
                    continue
                res = set(mul_points_list[i]).intersection(
                    set(mul_points_list[j]))
                if len(res) != 0:
                    # Remove i and j entries
                    if self.rect_list[i].name in ids_list:
                        ids_list.remove(self.rect_list[i].name)
                    if self.rect_list[j].name in ids_list:
                        ids_list.remove(self.rect_list[j].name)
        return ids_list


if __name__ == "__main__":
    ip_file = "trial_ip.txt"
    ip_file = "ip.txt"
    with open(ip_file) as fp:
        contents = fp.read()
    ip_entries = [ip for ip in contents.split("\n") if not ip.strip() == ""]

    rect_list = []

    for ip in ip_entries:
        rect = Rect.get_rect_from_entry(ip)
        rect_list.append(rect)

    manager = RectManager(rect_list)

    # part-1
    print("Number of overlaps -")
    overlaps = manager.find_overlaps()
    print(len(overlaps))

    # part-2
    print("Unique ID")
    unique_id = manager.find_no_overlap_id()
    print(unique_id[0])
