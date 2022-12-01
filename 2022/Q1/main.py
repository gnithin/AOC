if __name__ == "__main__":
    elves_calories = []
    with open("ip2.txt", "r") as fp:
        calories = 0
        for line in fp:
            l = line.strip()
            if l == "":
                elves_calories.append(calories)
                calories = 0
                continue
            calories += int(l)
        if calories > 0:
            elves_calories.append(calories)
    print(sum(sorted(elves_calories)[-1:-4:-1]))
