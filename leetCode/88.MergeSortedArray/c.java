class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        for (int ai : nums2) {
            insertEmplement(ai, nums1, m);
            m++;
        }
    }

    public static void insertEmplement(int x, int[] nums1, int m) {
        boolean found = false;
        for (int i = 0; i < m; i++) {
            if (nums1[i] > x) {
                found = true;
                // tìm thấy, dịch các phần tử từ i đến m - 1 sang phải 1 vị trí, dịch từ cuối
                for (int k = m - 1; k >= i; k--) {
                    nums1[k + 1] = nums1[k];
                }
                // Thêm x vào vị trí i
                nums1[i] = x;
                break;

            }
        }
        // Nếu không tìm thấy số nào lớn hơn x thì x sẽ được thêm vào cuối mảng
        if (found == false) {
            nums1[m] = x;
        }
    }
}