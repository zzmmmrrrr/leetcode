
"""
注意这里的dfs遇到水域会返回，也就是说dfs只能
"""
def find_lands(grid: list[list[str]]) -> int:
    m, n = len(grid), len(grid[0])  # m行n列
    dx = [1, -1, 0, 0]
    dy = [0, 0, 1, -1]

    def dfs(i, j):
        if i < 0 or i >= m or j < 0 or j >= n or grid[i][j] != "1":
            return
        grid[i][j] = "2"
        for k in range(4):
            dfs(i + dx[k], j + dy[k])

    ans = 0
    for i in range(m):
        for j in range(n):
            if grid[i][j] == "1":
                dfs(i, j)
                ans += 1

    return ans


grid = [
    ["1", "1", "1", "1"],
    ["1", "0", "0", "1"],
    ["0", "1", "0", "0"],
    ["1", "0", "1", "0"],
]

ans = find_lands(grid)
print(ans)
