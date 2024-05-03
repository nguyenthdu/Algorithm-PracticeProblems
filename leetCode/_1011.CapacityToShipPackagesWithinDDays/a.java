package leetCode._1011.CapacityToShipPackagesWithinDDays;

class Solution {
    public int shipWithinDays(int[] weights, int days) {
        int left = 0;
        int right = 0;
        for (int w : weights) {
            left = Math.max(left, w);
            right += w;
        }
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (isValidTime(weights, days, mid)) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        return left;
    }

    public boolean isValidTime(int[] weights, int days, int mid) {
        int totalDay = 1;
        int curCap = 0;
        for (int w : weights) {
            curCap += w;
            if (curCap > mid) {
                totalDay++;
                curCap = w;
            }
        }
        return totalDay <= days;
    }
}