def merge(nums1: list[int], nums2: list[int]) -> list[int]:
    ans = []
    i, j = 0, 0
    len1, len2 = len(nums1), len(nums2)
    while i < len1 and j < len2:
        if nums1[i] < nums2[j]:
            ans.append(nums1[i])
            i += 1
        else:
            ans.append(nums2[j])
            j += 1
    while i < len1:
        ans.append(nums1[i])
        i += 1

    while j < len2:
        ans.append(nums2[j])
        j += 1

    return ans


nums1 = [1, 2, 3, 4, 5, 6]
nums2 = [-2, 3, 323, 434343, 9999999]
print(merge(nums1, nums2))
