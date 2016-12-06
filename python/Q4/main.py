import sys
import re
from pprint import pprint


def order_names(name, reqd_len):
    char_dict = {}
    for char in name:
        char_dict[char] = char_dict.get(char, 0) + 1

    # reverse clubbing
    freq_dict = {}
    for char, freq in char_dict.items():
        if freq in freq_dict:
            freq_dict[freq].append(char)
            freq_dict[freq] = sorted(freq_dict[freq])
        else:
            freq_dict[freq] = [char]

    return ''.join(
        [
            ''.join(map(str, char_list))
            for _, char_list in sorted(freq_dict.items(), reverse=True)
        ]
    )[:reqd_len]


def get_valid_names(names_list):
    all_items = []
    for item in names_list:
        sector_id = item['sectorId']
        name = item['original_name']

        sentence = get_valid_sentence(name, sector_id)

        all_items.append({
            "name": sentence,
            "sectorId": sector_id
        })
    return all_items


def get_valid_sentence(sentence, sector_id):
    final_list = []
    for word in sentence.split('-'):
        new_word = ''
        for char in word:
            offset = ord(char) - ord('a')
            offset += sector_id
            offset = (offset % 26) + ord('a')
            new_word += chr(offset)
        final_list.append(new_word)
    return ' '.join(final_list)

if __name__ == "__main__":
    regex = re.compile(r'([\w-]+)\-(\d+)\[(\w+)]')
    ip_list = [
        {
            'name': order_names(
                reMatch.group(1).replace('-', ''),
                len(reMatch.group(3))
            ),
            'sectorId': int(reMatch.group(2)),
            'checksum': reMatch.group(3),
            'original_name': reMatch.group(1)
        }
        for ip in sys.stdin for reMatch in regex.finditer(ip.strip())
    ]

    valid_items = filter(
        lambda item: item['name'] == item['checksum'],
        ip_list
    )

    valid_count = sum([item['sectorId'] for item in valid_items])

    print(valid_count)
    print(len(valid_items))

    # Part 2 for the question
    names_and_sectors = get_valid_names(valid_items)

    # search for names with north on them
    for item in names_and_sectors:
        if 'north' in item['name']:
            print(item)

