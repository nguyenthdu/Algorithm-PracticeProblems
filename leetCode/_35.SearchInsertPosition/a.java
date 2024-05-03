package leetCode._35.SearchInsertPosition;

class Solution {
    public int searchInsert(int[] nums, int target) {
        int index = -1;
        int left = 0;
        int right = nums.length - 1;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) {
                index = mid;
            }
            if (nums[mid] > target) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        if (index == -1) {
            if (nums[left] < target) {
                return left + 1;
            } else {
                return left;
            }
        } else {
            return index;
        }

    }

}