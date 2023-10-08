class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        # for i in range(n):
        #     self.insertElement(nums2[i], nums1, m)
        #     m += 1
        for i in nums2:
            self.insertElement(i, nums1, m)
            m += 1

    def insertElement(self, x, arr, m):
        existing = False
        for k in range(m):
            if arr[k] > x:
                existing = True
                for j in range(m, k, -1):
                    arr[j] = arr[j-1]
                # arr[k+1:m+1] = arr[k:m]
                arr[k] = x
                break
        if not existing:
            arr[m] = x


# using Complexity
def merge(self, nums1, m, nums2, n):
    while m > 0 and n > 0:

        if nums1[m-1] >= nums2[n-1]:
            # giai thich: neu nums1[m-1] >= nums2[n-1] thi gan nums1[m+n-1] = nums1[m-1] va
            # -> dem phan tu cuoi cung cua nums1 len phan tu cuoi cung cua nums1 + nums2 va do dai cua nums1 n+m-1
            nums1[m+n-1] = nums1[m-1]
            # giai thich: neu n > 0 thi n la so phan tu con lai cua nums2, nen gan nums1[:n] = nums2[:n] la gan phan tu con lai cua nums2 vao nums1
            m -= 1
        else:
            # giai thich: neu nums1[m-1] < nums2[n-1] thi gan nums1[m+n-1] = nums2[n-1] va n -= 1
            # -> dem phan tu cuoi cung cua nums2 len phan tu cuoi cung cua nums1 + nums2 va do dai cua nums2 n+m-1
            nums1[m+n-1] = nums2[n-1]
            n -= 1
    if n > 0:
        # giai thich: neu n > 0 thi n la so phan tu con lai cua nums2, nen gan nums1[:n] = nums2[:n] la gan phan tu con lai cua nums2 vao nums1
        nums1[:n] = nums2[:n]
