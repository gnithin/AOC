import sys
import re


def find_num_ABBA_ip(ipList):
    search_regex = r'((\w)((?!\2)\w))\3\2'
    split_regex = r'\[(.*?)\]'
    count = 0

    for ip in ipList:
        all_vals = re.split(split_regex, ip)
        str_condn = (lambda x: len(re.findall(search_regex, x)) > 0)

        matched_li = (
            [
                str_condn(should_pass)
                for should_pass in all_vals[0::2]
            ]
        )

        unmatched_li = (
            [
                not str_condn(shouldnt_pass)
                for shouldnt_pass in all_vals[1::2]
            ]
        )

        if any(matched_li) and all(unmatched_li):
            count += 1
    return count


def find_num_ssl_ip(ipList):
    split_regex = r'\[(.*?)\]'
    count = 0

    for ip in ipList:
        all_vals = re.split(split_regex, ip)
        str_condn = (lambda x: len(re.findall(search_regex, x)) > 0)

        outside_vals = all_vals[0::2]
        inside_vals = all_vals[1::2]


if __name__ == "__main__":
    ipList = [ip.strip() for ip in sys.stdin]

    num_valid_ip = find_num_ABBA_ip(ipList)
    print(num_valid_ip)

    num_SSL_ip = find_num_ssl_ip(ipList)
    print(num_SSL_ip)
