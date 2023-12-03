public class Main {

    public static void main(String[] args) {
        int[] nums = { 12, 345, 2, 6, 7896 };
        System.out.println(findNumbers(nums));
    }

    public static int findNumbers(int[] nums) {
        int count = 0;
        for (int i = 0; i < nums.length; i++) {
            int n = 0;
            while (nums[i] > 0) {
                nums[i] = nums[i] / 10;
                n++;

            }
            if (n % 2 == 0) {
                count++;
            }
        }
        return count;
    }

}