import hashlib
import re
from dataclasses import dataclass, field
from typing import List, Set


@dataclass
class Point:
    x: int = field(default=0)
    y: int = field(default=0)
    z: int = field(default=0)

    def __str__(self):
        return f"{self.x},{self.y},{self.z}"


@dataclass
class Moons:
    positions: List[Point]
    velocities: List[Point]
    prev_states: Set[str] = field(default_factory=set)

    def __str__(self):
        entries = []
        for i in range(len(self.positions)):
            p = self.positions[i]
            v = self.velocities[i]
            entries.append(
                f"position: <x:{p.x}, y:{p.y}, z:{p.z}>\t"
                f"velocity: <x:{v.x}, y:{v.y}, z:{v.z}>"
            )
        return "\n".join(entries)

    def perform_steps(self, steps: int):
        for _ in range(steps):
            self.perform_step()

    @staticmethod
    def get_velocity(p1, p2):
        diff = p1 - p2
        if diff > 0:
            return -1
        elif diff < 0:
            return 1
        return 0

    def perform_step(self):
        num_entries = len(self.velocities)

        velocity_diff_list = []
        for i in range(num_entries):
            velocity_diff_list.append(Point())

        for i in range(num_entries):
            curr_position = self.positions[i]
            velocity_diff = velocity_diff_list[i]
            for j in range(num_entries):
                # Skip the same entry
                if i == j:
                    continue
                p = self.positions[j]
                velocity_diff.x += self.get_velocity(curr_position.x, p.x)
                velocity_diff.y += self.get_velocity(curr_position.y, p.y)
                velocity_diff.z += self.get_velocity(curr_position.z, p.z)

        state_list = []
        for i in range(num_entries):
            velocity_diff = velocity_diff_list[i]
            # Update velocity
            self.velocities[i].x += velocity_diff.x
            self.velocities[i].y += velocity_diff.y
            self.velocities[i].z += velocity_diff.z

            # Update position
            self.positions[i].x += self.velocities[i].x
            self.positions[i].y += self.velocities[i].y
            self.positions[i].z += self.velocities[i].z

            # Store the value of the entries in a hash
            state_list.append(self.velocities[i].x)
            state_list.append(self.velocities[i].y)
            state_list.append(self.velocities[i].z)
            state_list.append(self.positions[i].x)
            state_list.append(self.positions[i].y)
            state_list.append(self.positions[i].z)

        state_str = ",".join(map(str, state_list))
        hash_res = hashlib.md5(state_str.encode())
        hash = hash_res.hexdigest()
        return hash

    def get_energy(self) -> int:
        individual_energies = map(
            lambda x: x[0] * x[1],
            zip(self.get_potential_energy(), self.get_kinetic_energy())
        )
        return sum(individual_energies)

    def get_potential_energy(self):
        return [abs(m.x) + abs(m.y) + abs(m.z) for m in self.positions]

    def get_kinetic_energy(self):
        return [abs(m.x) + abs(m.y) + abs(m.z) for m in self.velocities]

    def find_recurring_step(self) -> int:
        num_planets = len(self.velocities)

        recurring_steps = []
        # Todo add for other axes
        axes = ["x"]  # , "y", "z"]
        for axis in axes:
            step = 0
            axis_period = [0 for _ in range(num_planets)]
            prev_states: List[Set[str]] = [set() for _ in range(num_planets)]

            planet_hashes = []
            for i in range(num_planets):
                v = self.velocities[i].__dict__[axis]
                p = self.positions[i].__dict__[axis]
                hash_val = self.get_hash(v, p)
                prev_states[i].add(hash_val)

            planet_hashes = self.perform_step_for_axis(axis)
            for i in range(num_planets):
                if axis_period[i] != 0:
                    continue
                hash = planet_hashes[i]
                planet_set = prev_states[i]
                if hash in planet_set:
                    axis_period[i] = step

            while 0 in axis_period:
                step += 1
                for i, h in enumerate(planet_hashes):
                    prev_states[i].add(h)
                planet_hashes = self.perform_step_for_axis(axis)
                for i in range(num_planets):
                    if axis_period[i] != 0:
                        continue
                    hash = planet_hashes[i]
                    planet_set = prev_states[i]
                    if hash in planet_set:
                        if i == 1:
                            print(f"Found - {hash} in {prev_states[i]}")
                            print(f"Step - {step}")
                        axis_period[i] = step

            # Find the LCM of the axis steps
            print("Axis - " + axis)
            recurring_steps.append(self.find_LCM(*axis_period))
        print(f"Steps - {recurring_steps}")
        return 0

    def perform_step_for_axis(self, axis):
        num_entries = len(self.velocities)
        hash_list = []

        velocity_diff_list = []
        for i in range(num_entries):
            velocity_diff_list.append(Point())

        for i in range(num_entries):
            curr_position = self.positions[i]
            velocity_diff = velocity_diff_list[i]
            for j in range(num_entries):
                # Skip the same entry
                if i == j:
                    continue
                p = self.positions[j]
                velocity_diff.__dict__[axis] += self.get_velocity(
                    curr_position.__dict__[axis],
                    p.__dict__[axis]
                )

        for i in range(num_entries):
            velocity_diff = velocity_diff_list[i]
            # Update velocity
            self.velocities[i].__dict__[axis] += velocity_diff.__dict__[axis]

            # Update position
            self.positions[i].__dict__[axis] += self.velocities[i].__dict__[axis]

            # Store the value of the entries in a hash
            v = self.velocities[i].__dict__[axis]
            p = self.positions[i].__dict__[axis]
            hash_val = self.get_hash(v, p)
            if i == 1:
                print(f"{v},{p} -> {hash_val}")
            hash_list.append(hash_val)

        return hash_list

    def get_hash(self, v, p):
        hash_str = f"{v},{p}"
        hash_res = hashlib.md5(hash_str.encode())
        return hash_res.hexdigest()

    def find_LCM(self, *args):
        print(f"LCM - {args}")
        return 0


if __name__ == "__main__":
    initial_positions = []
    initial_velocities = []
    filename = "ip1.txt"
    # filename = "ip2.txt"
    with open(filename, "r") as fp:
        for line in fp:
            parts = re.findall(r'[xyz]=(-?\d+)', line)
            initial_positions.append(Point(*map(int, parts)))
            initial_velocities.append(Point())

    # Moon steps
    moons = Moons(initial_positions, initial_velocities)
    # steps = 4686774924
    # for i in range(steps):
    #     moons.perform_steps(1)
    #     if i % 10000:
    #         print(i)
    #     # print(f"Step - {i + 1} {'*' * 20} ")
    #     # print(moons)
    # print(f"Energy - {moons.get_energy()}")
    print(f"Step - {moons.find_recurring_step()}")
