# t = 0

# for line in open("./data.txt"):
#     a, b, x, y = map(int, line.replace(",", "-").split("-"))
#     if a <= x and b >= y or x <= a and y >= b:
#         t += 1

# print(t)

t = 0

for line in open("./data.txt"):
    a, b, x, y = map(int, line.replace(",", "-").split("-"))
    if set(range(a, b + 1)) & set(range(x, y + 1)):
        t += 1

print(t)