import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Scanner;

public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        /* 
         * 
         */
        int i = 1;

        while (sc.hasNext()) {
            String s = sc.nextLine();
            System.out.println(i + " " + s);
            i++;
        }

        // int n = sc.nextInt();
        // int count = 0;
        // do {
        // long x;
        // try {
        // x = sc.nextLong();
        // System.out.println(x + " can be fitted in:");
        // } catch (Exception e) {
        // System.out.println(sc.next() + " can't be fitted anywhere.");
        // count++;
        // continue;
        // }

        // if (x >= Byte.MIN_VALUE && x <= Byte.MAX_VALUE) {
        // System.out.println("* byte");
        // }
        // if (x >= Short.MIN_VALUE && x <= Short.MAX_VALUE) {
        // System.out.println("* short");
        // }
        // if (x >= Integer.MIN_VALUE && x <= Integer.MAX_VALUE) {
        // System.out.println("* int");
        // }
        // if (x >= Long.MIN_VALUE && x <= Long.MAX_VALUE) {
        // System.out.println("* long");
        // }
        // count++;

        // } while (count < n);
    }

}
