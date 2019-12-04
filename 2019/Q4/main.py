def match_criteria(password):
    if len(str(password)) != 6:
        return False
    entries = [int(i) for i in str(password)]
    found_zero = False
    zeroes = []
    for i in range(len(entries) - 1):
        diff = entries[i] - entries[i + 1]
        if diff == 0:
            found_zero = True
            zeroes.append(True)
        else:
            zeroes.append(False)
        if diff > 0:
            return False

    if not found_zero:
        return False

    # Find true between false
    for i in range(len(zeroes)):
        if not zeroes[i]:
            continue
        if i > 0 and (i + 1) < (len(entries) - 1):
            if zeroes[i - 1] == False and zeroes[i + 1] == False:
                return True
        elif i == 0:
            if zeroes[i + 1] == False:
                return True
        else:
            if zeroes[i - 1] == False:
                return True
    return False


if __name__ == "__main__":
    c = 0
    for i in range(356261, 846304):
        if match_criteria(i):
            c += 1
    print(c)
