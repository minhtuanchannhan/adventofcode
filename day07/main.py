# cwd = root = {}
# stack = []

# for line in open("./data.txt"):
#     line = line.strip()
#     if line[0] == "$":
#         if line[2] == "c":
#             dir = line[5:]
#             if dir == "/":
#                 cwd = root
#                 stack = []
#             elif dir == "..":
#                 cwd = stack.pop()
#             else:
#                 if dir not in cwd:
#                     cwd[dir] = {}
#                 stack.append(cwd)
#                 cwd = cwd[dir]
#     else:
#         x, y = line.split()
#         if x == "dir":
#             if y not in cwd:
#                 cwd[y] = {}
#         else:
#             cwd[y] = int(x)

# def solve(dir = root):
#     if type(dir) == int:
#         return (dir, 0)
#     size = 0
#     ans = 0
#     for child in dir.values():
#         s, a = solve(child)
#         size += s
#         ans += a
#     if size <= 100000:
#         ans += size
#     return (size, ans)

# print(solve()[1])

cwd = root = {}
stack = []

for line in open("./data.txt"):
    line = line.strip()
    if line[0] == "$":
        if line[2] == "c":
            dir = line[5:]
            if dir == "/":
                cwd = root
                stack = []
            elif dir == "..":
                cwd = stack.pop()
            else:
                if dir not in cwd:
                    cwd[dir] = {}
                stack.append(cwd)
                cwd = cwd[dir]
    else:
        x, y = line.split()
        if x == "dir":
            if y not in cwd:
                cwd[y] = {}
        else:
            cwd[y] = int(x)

def size(dir = root):
    if type(dir) == int:
        return dir
    return sum(map(size, dir.values()))

t = size() - 40_000_000

def solve(dir = root):
    ans = float("inf")
    if size(dir) >= t:
        ans = size(dir)
    for child in dir.values():
        if type(child) == int:
            continue
        q = solve(child)
        ans = min(ans, q)
    return ans

print(solve())