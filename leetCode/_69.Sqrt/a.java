
class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.mySqrt(8));
    }

    public int mySqrt(int x) {
        if (x < 2)
            return x;
        int left = 0;
        int right = x;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (x / mid < mid) {
                right = mid;
            } else {
                left = mid + 1;
            }

        }
        return left - 1;
    }
}