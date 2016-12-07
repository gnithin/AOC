import hashlib


def get_password_for_door_id(door_id, part1=True):
    password = ''
    if not part1:
        password = [1, 2, 3, 4, 5, 6, 7, 8]

    i = 0
    for password_num in range(8):
        char_found = False
        while not char_found:
            md5_func = hashlib.md5()
            md5_func.update(door_id + str(i))
            hex_val = md5_func.hexdigest()
            i += 1
            if hex_val[:5] == '00000':
                if part1:
                    password += hex_val[5]
                    print("Found - ", hex_val[5])
                    break
                else:
                    pos = hex_val[5]
                    char = hex_val[6]
                    if pos in '01234567' and type(password[int(pos)]) == int:
                        password[int(pos)] = char
                        print("Found- ", char, pos)
                        break
    if not part1:
        return ''.join(password)
    return password


if __name__ == "__main__":
    ip_text = 'ffykfhsq'
    # ip_text = 'abc'

    # password = get_password_for_door_id(ip_text)
    # print(password)

    new_password = get_password_for_door_id(ip_text, part1=False)
    print(new_password)
