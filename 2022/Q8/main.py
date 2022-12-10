from functools import reduce


def get_num_visible_trees(matrix):
    visible_indices = set()
    num_rows = len(matrix)
    num_cols = len(matrix[0])

    for r in range(1, num_rows - 1):
        for c in range(1, num_cols - 1):
            curr_val = matrix[r][c]

            # Go vertically and horizontally from that index
            if any([matrix[i][c] >= curr_val for i in range(r - 1, -1, -1)]) and \
                    any([matrix[j][c] >= curr_val for j in range(r + 1, num_rows)]) and \
                    any([matrix[r][k] >= curr_val for k in range(c - 1, -1, -1)]) and \
                    any([matrix[r][l] >= curr_val for l in range(c + 1, num_cols)]):
                pass
            else:
                visible_indices.add((r, c))

    # print(visible_indices)
    return len(visible_indices) + ((2 * num_rows) + (2 * num_cols) - 4)


def get_highest_scenic_score(matrix):
    scenic_scores = []
    num_rows = len(matrix)
    num_cols = len(matrix[0])

    for r in range(1, num_rows - 1):
        for c in range(1, num_cols - 1):
            curr_val = matrix[r][c]

            scores = []
            score = 0
            for i in range(r - 1, -1, -1):
                score += 1
                if curr_val <= matrix[i][c]:
                    break
            scores.append(score)

            score = 0
            for j in range(r + 1, num_rows):
                score += 1
                if curr_val <= matrix[j][c]:
                    break
            scores.append(score)

            score = 0
            for k in range(c - 1, -1, -1):
                score += 1
                if curr_val <= matrix[r][k]:
                    break
            scores.append(score)

            score = 0
            for l in range(c + 1, num_cols):
                score += 1
                if curr_val <= matrix[r][l]:
                    break
            scores.append(score)
            print(scores)

            scenic_scores.append(reduce(lambda x, y: x * y, scores))

    print(scenic_scores)
    return max(scenic_scores)


if __name__ == "__main__":
    matrix = []
    with open("ip2.txt", "r") as fp:
        for line in fp:
            row = []
            for ch in line.strip():
                row.append(int(ch))
            matrix.append(row)
    # print(matrix)

    # num_visible = get_num_visible_trees(matrix)
    # print(num_visible)
    print(get_highest_scenic_score(matrix))
