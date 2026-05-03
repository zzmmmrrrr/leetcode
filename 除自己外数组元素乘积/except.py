def productExceptSelf(nums: list[int]) -> list[int]:
    n = len(nums)
    ans = [1] * n
    # i from 1 to n-1
    for i in range(1, n):
        ans[i] = ans[i - 1] * nums[i - 1]

    # right表示当前nums[i]右侧元素乘积
    right = 1
    for i in range(n - 1, -1, -1):
        ans[i] *= right
        right *= nums[i]

    return ans


nums = [1, 2, 3, 4, 5, 6]
print(productExceptSelf(nums))
