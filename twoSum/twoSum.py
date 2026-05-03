def two_sum(nums: list[int], target: int) -> list[int]:
    hash = {}
    # hash[num]=i,存储元素对应的角标
    for i, num in enumerate(nums):
        diff = target - num
        if diff in hash:
            return [hash[diff], i]
        hash[num] = i

    return []


test = [1, -1, 2, 3, 4, 5, 6, -10, 23, 444]
target = 0

ans = two_sum(test, target)
print(ans)
