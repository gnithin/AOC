class LayerManager:
    def __init__(self, layers):
        self.layers = layers

    def get_part_1(self):
        zero_layers = list(map(lambda x: (x, x.find_num_digits(0)), self.layers))
        layer = min(zero_layers, key=lambda x: x[1])[0]
        print(layer)
        return layer.find_num_digits(1) * layer.find_num_digits(2)


class Layer:
    def __init__(self, w, h, ip):
        self.matrix = []
        self.w = w
        self.h = h
        for i in range(w):
            self.matrix.append([0 for _ in range(h)])

        e = 0
        for i in range(w):
            for j in range(h):
                self.matrix[i][j] = ip[e]
                e += 1

    def __str__(self):
        return str(self.matrix)

    def find_num_digits(self, digit):
        count = 0
        for i in range(self.w):
            for j in range(self.h):
                if self.matrix[i][j] == digit:
                    count += 1
        return count


def get_layers(w, h, ip_list):
    layers = []
    size = w * h
    i = 0
    while i < len(ip_list):
        layers.append(Layer(w, h, ip_list[i:(i + size)]))
        i += size
    return layers


if __name__ == "__main__":
    # file_name = "ip1.txt"
    # width = 3
    # height = 2

    file_name = "ip2.txt"
    width = 25
    height = 6

    with open(file_name, "r") as fp:
        ip_list = list(map(int, [c for c in fp.read().strip()]))

    layers = get_layers(height, width, ip_list)
    mgr = LayerManager(layers)
    layer = mgr.get_part_1()
    print(layer)
