class Solution:
    def removeDuplicate(__self__, nums):
        i = 0
        while i < len(nums)-1:
            if nums[i] in nums[i+1:]:
                nums.pop(i)
            else:
                i += 1
        return len(nums)


print(Solution().removeDuplicate([1, 1, 2]))
