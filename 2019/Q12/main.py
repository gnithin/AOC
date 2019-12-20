import re
from dataclasses import dataclass, field
from typing import List


@dataclass
class Point:
    x: int = field(default=0)
    y: int = field(default=0)
    z: int = field(default=0)


@dataclass
class Moons:
    positions: List[Point]
    velocities: List[Point]

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

    def perform_step(self):
        num_entries = len(self.velocities)

        def get_velocity(p1, p2):
            diff = p1 - p2
            if diff > 0:
                return -1
            elif diff < 0:
                return 1
            return 0

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
                velocity_diff.x += get_velocity(curr_position.x, p.x)
                velocity_diff.y += get_velocity(curr_position.y, p.y)
                velocity_diff.z += get_velocity(curr_position.z, p.z)

        for i in range(num_entries):
            curr_position = self.velocities[i]
            velocity_diff = velocity_diff_list[i]
            # Update velocity
            self.velocities[i].x += velocity_diff.x
            self.velocities[i].y += velocity_diff.y
            self.velocities[i].z += velocity_diff.z

            # Update position
            self.positions[i].x += self.velocities[i].x
            self.positions[i].y += self.velocities[i].y
            self.positions[i].z += self.velocities[i].z

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
    steps = 4686774924
    for i in range(steps):
        moons.perform_steps(1)
        if i % 10000:
            print(i)
        # print(f"Step - {i + 1} {'*' * 20} ")
        # print(moons)
    print(f"Energy - {moons.get_energy()}")
