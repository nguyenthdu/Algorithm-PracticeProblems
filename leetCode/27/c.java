class Solution {
    public int removeElement(int[] nums, int val) {

        for (int i = 0; i < nums.length; i++) {
            if (i == val) {
                for (int j = i + 1; j < nums.length; j++) {
                    nums[i] = nums[j];
                }
            }
        }
        return nums.length;
    }
}