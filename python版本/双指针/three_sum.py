def three_sum(nums: list[int]) -> list[int]:
    # 1.记录数组长度，并对nums排序
    n = len(nums)
    nums.sort()
    ans = []

    # 2.固定第一个数字，角标从0遍历到n-3
    # 原因是避免越界，最后一个三元数组角标是n-3,n-2,n-1
    for i in range(n - 2):
        # 3.剪枝操作，大数据量时有明显优化,可以先只记住第一个操作
        # 如果仅第一个元素就大于0，直接break

        # 如果固定nums[i]时最小的三元组已经大于0，
        # 由于nums[i]后续元素更大，因此全局都不可能再有合法解
        # 如果固定nums[i]前提下，最大的三元组已经小于0，
        # 代表nums[i]做首元素时不可能有合法解，换更大的首元素
        if nums[i] > 0:
            break
        if nums[i] + nums[i + 1] + nums[i + 2] > 0:
            break
        if nums[i] + nums[n - 2] + nums[n - 1] < 0:
            continue

        # 4.去重，保证相同元素只在第一次出现时进行处理
        if i > 0 and nums[i] == nums[i - 1]:
            continue

        # 5.左右指针不断移动，左指针确定第二个元素，右指针确定第三个元素
        l, r = i + 1, n - 1
        while l < r:
            cur_sum = nums[i] + nums[l] + nums[r]

            # 对于每组合法解，都要跳过相同元素
            if cur_sum == 0:
                ans.append([nums[i], nums[l], nums[r]])
                while l < r and nums[l] == nums[l + 1]:
                    l += 1
                l += 1
                while l < r and nums[r] == nums[r - 1]:
                    r -= 1
                r -= 1
            elif cur_sum < 0:
                l += 1
            else:
                r -= 1
    return ans


if __name__ == "__main__":
    nums = [1, -2, 3, 4, -5, -5, 2, 2, 2, 2, -1, 1, -3, -4, 5, 5, 5]
    result = three_sum(nums)
    for row in result:
        print(" ".join(map(str, row)))
