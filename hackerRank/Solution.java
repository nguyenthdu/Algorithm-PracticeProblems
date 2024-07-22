import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Iterator;
import java.util.List;
import java.util.Scanner;
import java.util.Set;

public class Solution {

    public static void main(String[] args) {

        List<List<Integer>> list = new ArrayList<List<Integer>>();
        int n = list.size();
        int result = 0;
        // tính tổng 2 đường chéo ma trận
        for (int i = 0; i < n; i++) {
            result += list.get(i).get(i) - list.get(i).get(n - i - 1);
        }
        System.out.println(Math.abs(result));
    }

    public static int diagonalDifference(List<List<Integer>> arr) {
        int n = arr.size();
        int result = 0;
        for (int i = 0; i < n; i++) {
            result += arr.get(i).get(i) - arr.get(i).get(n - i - 1);

        }
        return Math.abs(result);

    }

    public static void plusMinus(List<Integer> arr) {
        // Write your code here
        float positive = 0;
        float negative = 0;
        float zeros = 0;
        int n = arr.size();
        for (Integer a : arr) {
            if (a > 0) {
                positive++;
            } else if (a == 0) {
                zeros++;
            } else {
                negative++;
            }
        }
        System.out.println(positive / n);
        System.out.println(negative / n);
        System.out.println(zeros / n);

    }

    public static void staircase(int n) {
        // Write your code here
        for (int i = n; i > 0; i--) {
            for (int j = 1; j <= n; j++) {
                if (j < i) {
                    System.out.print(" ");
                } else {
                    System.out.print("#");
                }

            }
            System.out.println();
        }
    }

    public static void miniMaxSum(List<Integer> arr) {
        // Write your code here
        long min = 0;
        long max = 0;
        long sum = 0;
        for (int i = 0; i < arr.size(); i++) {
            sum += arr.get(i);
        }
        min = sum - arr.get(0);
        max = sum - arr.get(0);
        for (int i = 1; i < arr.size(); i++) {
            if (min > sum - arr.get(i)) {
                min = sum - arr.get(i);
            }
            if (max < sum - arr.get(i)) {
                max = sum - arr.get(i);
            }
        }
        System.out.println(min + " " + max);
    }

    public class Solution {
        public static void main(String[] args) {

            Scanner in = new Scanner(System.in);
            int testCases = Integer.parseInt(in.nextLine());
            while (testCases > 0) {
                String line = in.nextLine();

                // Write your code here
                boolean matchFound = false;
                String regex = "<(.+)>([^<]+)</\\1>";
                while (line.matches(".*" + regex + ".*")) {
                    line = line.replaceAll(regex, "$2");
                    matchFound = true;
                }
                if (matchFound) {
                    System.out.println(line);

                } else {
                    System.out.println("None");
                }

                testCases--;
            }
        }
    }

    public void CalcBigInterger() {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        sc.close();
        java.math.BigInteger aBig = new java.math.BigInteger(a);
        java.math.BigInteger bBig = new java.math.BigInteger(b);
        System.out.println(aBig.add(bBig));
        System.out.println(aBig.multiply(bBig));
    }

    public void handleSortBigDecima() {
        // Input
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        String[] s = new String[n + 2];
        for (int i = 0; i < n; i++) {
            s[i] = sc.next();
        }
        sc.close();
        // sort them in descending order
        java.util.Arrays.sort(s, 0, n, new java.util.Comparator<String>() {
            @Override
            public int compare(String a1, String a2) {
                java.math.BigDecimal a = new java.math.BigDecimal(a1);
                java.math.BigDecimal b = new java.math.BigDecimal(a2);
                return b.compareTo(a);
            }
        });
        for (BigDecimal num : s) {
            System.out.println(num);
        }
    }

    public void isProbablePrime() {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        String n = bufferedReader.readLine();

        bufferedReader.close();

        BigInteger bigInteger = new BigInteger(n);
        if (bigInteger.isProbablePrime(1)) {
            System.out.println("prime");
        } else {
            System.out.println("not prime");
        }
    }

    public static void MyCalculator() {
        Scanner in = new Scanner(System.in);
        int n = 0;
        int p = 0;
        while (in.hasNextInt()) {
            n = in.nextInt();
            p = in.nextInt();
            try {
                if (n < 0 || p < 0) {
                    throw new Exception("n or p should not be negative.");
                } else if (n == 0 && p == 0) {
                    throw new Exception("n and p should not be zero.");
                } else {
                    System.out.println((int) Math.pow(n, p));
                }
            } catch (Exception e) {
                System.out.println(e);
            }

        }

    }

    public void test() {
        ArrayList<Integer> list = new ArrayList<Integer>();
        list.add(1);
        list.add(2);
        list.add(3);
        list.add(4);
        list.add(5);
        plusMinus(list);
        list.forEach((i,v)->{
            System.out.println(i);
        });
        };
    }

}