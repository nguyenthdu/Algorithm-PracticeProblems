import java.io.*;
import java.util.*;;

class BinarySearch {
    public static void main(String[] args) {
        int[] list = { 1, 2, 3, 4, 5, 6, 7, 8 };
        int key = 5;
        System.out.println(BSearch(list, key));

    }

    static int BSearch(int list[], int key) {

        Arrays.sort(list);
        int left = 0;
        int right = list.length - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (key == list[mid])
                return mid;
            if (key < list[mid])
                left = mid + 1;
            if (key > list[mid])
                right = mid - 1;
        }
        return -1;

    }
}