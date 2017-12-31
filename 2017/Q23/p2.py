from math import sqrt; from itertools import count, islice

def is_prime(n):
    return n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

b = 106500
c = 123501
h = 0

for b in range(b, c, 17):
	if not is_prime(b):h += 1

print(h)
