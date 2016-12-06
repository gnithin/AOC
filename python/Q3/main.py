import sys


def is_triangle(triad):
    triad.sort()
    return sum(triad[:2]) > triad[-1]


def get_num_triangles(tri_list):
    return len([1 for triad in tri_list if is_triangle(triad)])


if __name__ == "__main__":
    is_part_2 = True

    ip_list = [map(int, ip.strip().split()) for ip in sys.stdin]

    # For part 2 of the question
    if is_part_2:
        rotate_li = reduce(lambda x, y: x+y, zip(*ip_list))
        ip_list = [list(rotate_li[i:i+3]) for i in range(0, len(rotate_li), 3)]

    tri_count = get_num_triangles(ip_list)
    print("Number valid triangles - ", tri_count)
