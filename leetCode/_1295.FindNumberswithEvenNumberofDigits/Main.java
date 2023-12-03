
public class Main {
    public static int findNumber(int[] nums) {
        int count = 0;
        for (int i : nums) {
            int digit = countDigit(i);
            if (digit % 2 == 0) {
                count++;
            }

        }
        return count;
    }
    // Dem so chua so

    private static int countDigit(int i) {
        int count = 0;
        while (i > 0) {
            i = i / 10;
            count++;
        }
        return count;
    }

    public static void main(String[] args) {
        int[] nums = { 12, 345, 2, 6, 7896 };
        for (int i : nums) {
            System.out.println("ai: " + i + "");

        }
        int[] n2 = new int[4];
        n2[0] = 1;
        n2[1] = 2;
        n2[2] = 3;
        n2[3] = 4;
        for (int i = 0; i < n2.length; i++) {
            System.out.println("n2: " + n2[i] + "");
        }
    }
}
