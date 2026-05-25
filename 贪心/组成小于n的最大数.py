"""
思路概括：
1. 对 digits 进行正向排序并去重，使后面利用索引获取最大值更方便。
2. dfs(digits, n_digits, ans, idx) 返回 bool 值，表示从当前 idx 开始是否能凑出严格小于 n 的解：
   - 当 idx == n_len 意味着每一位都相等，由于要求严格小于，属于非法撞车，直接返回 False 触发回溯。
   - 当选了更小的数，后面直接填最大值，返回 True。
   - 当选了相等的数，递归看下游。若下游失败，在当前位继续尝试更小的数。
3. 若同位数构造失败，则退而求其次构造长度减 1 且全由 digits 最大值组成的数字。

复杂度：
- 时间：O(K log K + L * K)，K 为 digits 长度，L 为 n 的位数。
- 空间：O(L)，用于存储位数数组及递归栈空间。
"""

def int_to_digits(n: int) -> list[int]:
    if n == 0:
        return [0]
    res = []
    while n > 0:
        res.append(n % 10)
        n //= 10
    res.reverse()
    return res
    
    # 【列表推导式简化版本】：
    # return [int(d) for d in str(n)]


def digits_to_int(digits: list[int]) -> int:
    ans = 0
    for d in digits:
        ans = ans * 10 + d
    return ans

    # 【列表推导式与内建函数简化版本】：
    # import functools
    # return functools.reduce(lambda x, y: x * 10 + y, digits)
    # 或者通过字符串中转：
    # return int("".join([str(d) for d in digits]))


def unique(nums: list[int]) -> list[int]:
    n = len(nums)
    if n <= 1:
        return nums

    ans = [nums[0]]
    for i in range(1, n):
        if nums[i] != nums[i - 1]:
            ans.append(nums[i])
    return ans


def dfs(digits: list[int], n_digits: list[int], ans: list[int], idx: int) -> bool:
    n_len = len(n_digits)
    
    # 递归终止条件：一路踩着等号走到了最后一位，说明最终构造出的数等于 n。
    # 因为要求【严格小于】，等于 n 是不合法的，直接返回 False 触发回溯。
    # 如果以后题目改成了【小于等于】，这里直接改成 return True 即可。
    if idx == n_len:
        return False

    # 贪心：因为 digits 是升序排序的，我们从大到小（从后往前）尝试可选数字
    for i in range(len(digits) - 1, -1, -1):
        d = digits[i]

        # 情况 1：当前数字大于 n 对应位的数字，不能选，继续看更小的
        if d > n_digits[idx]:
            continue

        # 情况 2：当前数字小于 n 对应位的数字
        if d < n_digits[idx]:
            ans[idx] = d
            max_digit = digits[-1] # 升序排序下，最后一个就是最大值
            # 贪心确定：当前位一旦变小，后面所有位必须填最大值才能保证整体最大
            for j in range(idx + 1, n_len):
                ans[j] = max_digit
            return True # 已经完美补齐，直接宣布同位数构造成功

        # 情况 3：d == n_digits[idx]
        ans[idx] = d
        # 试探下一步：看看下游能不能帮我凑出一个合法的、严格小于 n 后续分支的解
        if dfs(digits, n_digits, ans, idx + 1):
            return True # 下游成功，我也成功，直接向上汇报

        # 如果下游返回 False（说明后面撞车变回 n，或者后面无数字可选导致回溯）
        # 循环不中断，继续走下一次循环，在当前位尝试更小的数字

    return False # 当前位尝试了所有数字都无法拼出合法解，返回失败让上一层去变小


def max_num_less_than_n(digits: list[int], n: int) -> int:
    # 1. 预处理：升序排序（方便配合 -1 索引拿最大值）、去重、拆解 n 为位数组
    digits.sort()
    digits = unique(digits)
    n_digits = int_to_digits(n)
    n_len = len(n_digits)
    ans = [0] * n_len

    # 2. 核心贪心：尝试构造一个与 n 位数相同的数字
    # 只要 dfs 返回 True，说明 ans 数组里已经填好了最优解，直接转成 int 返回
    if dfs(digits, n_digits, ans, 0):
        return digits_to_int(ans)

    # 3. 退而求其次：如果同位数构造失败（如 n=2111, digits=[2,9]）
    # 则构造一个比 n 少一位的最大数，每一位都填 digits 里的最大值
    max_d = digits[-1]
    y = 0
    for i in range(n_len - 1):
        y = y * 10 + max_d
    return y


# 测试运行
if __name__ == "__main__":
    test_cases = [
        {"digits": [1, 3, 5, 9], "n": 5432, "expect": 5399},
        {"digits": [1, 2, 9], "n": 2151, "expect": 2129},
        {"digits": [2, 9], "n": 2111, "expect": 999},
        {"digits": [4, 3, 5, 2], "n": 5432, "expect": 5425} # 你最初卡住的那个特殊 case
    ]

    for case in test_cases:
        # 传入 digits 的切片副本，防止多 case 运行之间互相污染排序
        got = max_num_less_than_n(case["digits"][:], case["n"])
        status = "✓" if got == case["expect"] else "✗"
        print(f"digits={case['digits']}, n={case['n']} -> 实际输出: {got} ({status})")