import java.util.HashMap;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        int[] nums = { 2, 7, 11, 15 };
        // int[] result = twoSum(nums, 9);
        int[] result = twoSum2(nums, 9);
        System.out.println(result[0] + " " + result[1]);
    }

    // brute force: O(n^2)
    public static int[] twoSum(int[] nums, int target) {
        int[] result = new int[2];
        int i = 0;
        int j = 0;
        boolean found = false;
        while (i < nums.length && !found) {
            j = i + 1;
            while (j < nums.length && !found) {
                if (nums[i] + nums[j] == target) {
                    found = true;
                    result[0] = i;
                    result[1] = j;
                }

            }
        }

        return result;
    }

    // hash map: O(n)
    public static int[] twoSum2(int[] nums, int target) {
        int[] result = new int[2];
        Map<Integer, Integer> map = new HashMap<>();

        // put all elements into map
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
        }

        // find the complement
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            // check if the complement is in the map
            if (map.containsKey(complement) && map.get(complement) != i) {
                result[0] = i;
                result[1] = map.get(complement);
            }
        }

        return result;
    }
}
