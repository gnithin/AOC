from typing import Dict


class PlanetsMap:
    def __init__(self):
        self.planet_map: Dict[str, str] = {}
        self.cache = {}

    def add_mapping(self, parent, child):
        self.planet_map[child] = parent
        # Invalidate cache
        self.cache = {}

    def find_total_orbits(self):
        orbits = 0
        for planet in self.planet_map:
            orbits += self.find_orbits(planet)
        return orbits

    def find_orbits(self, planet: str):
        if planet in self.cache:
            return self.cache[planet]

        ret_val = 0
        if planet in self.planet_map:
            ret_val = 1 + self.find_orbits(self.planet_map[planet])
        self.cache[planet] = ret_val
        return ret_val

    def find_all_parent(self, planet):
        parent = [planet]
        if planet in self.planet_map:
            parent.extend(self.find_all_parent(self.planet_map[planet]))
        return parent

    def find_path(self, src, dest):
        src_parents = self.find_all_parent(src)
        dest_parents = self.find_all_parent(dest)

        longer_list = src_parents
        shorter_list = dest_parents
        if len(shorter_list) > len(longer_list):
            longer_list = dest_parents
            shorter_list = src_parents

        diff_len = len(longer_list) - len(shorter_list)
        longer_index = diff_len
        shorter_index = 0

        common_parent = None
        while shorter_index < len(shorter_list):
            if shorter_list[shorter_index] == longer_list[longer_index]:
                common_parent = shorter_list[shorter_index]
                break
            shorter_index += 1
            longer_index += 1

        if common_parent is None:
            return None

        path = shorter_list[:shorter_index + 1] + longer_list[:longer_index][::-1]
        return path


if __name__ == "__main__":
    with open("ip1.txt", "r") as fp:
        ip_list = [line.strip().split(")") for line in fp]

    planets = PlanetsMap()
    for (parent, child) in ip_list:
        planets.add_mapping(parent, child)

    # print(planets.find_total_orbits())
    path = planets.find_path("YOU", "SAN")
    print(len(path) - 3)
