package practice;

import java.util.ArrayList;

/**
 * ReverseArray
 */
public class ReverseArray {
    public static void main(String[] args) {
        // int[] arr = { 1, 2, 3, 4, 5 };
        // // int[] reversed = reverseArrayExtraArray(arr);
        // reverseArrayTwoPointer(arr, 0, arr.length - 1);
        // printArray(arr);
        // System.out.println();
        // khai bao mang int 10 phan tu
        ArrayList<String> arr = new ArrayList<String>();
        String[] arr1 = new String[10];
        for (int i = 0; i < 10; i++) {
            System.out.println(arr1[i]);
        }
    }

    public static int[] reverseArrayExtraArray(int[] arr) {
        int[] reversed = new int[arr.length];
        for (int i = 0; i < arr.length; i++) {
            reversed[i] = arr[arr.length - 1 - i];
        }
        return reversed;
    }

    public static void reverseArrayTwoPointer(int[] arr, int start, int end) {
        int temp;
        while (start < end) {
            temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }
    }

    public static void printArray(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            System.out.print(arr[i] + " ");
        }
    }

}
