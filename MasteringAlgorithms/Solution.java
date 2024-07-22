
public class Solution {
    public static void main(String[] args) {
        Hanoi(3, 'A', 'C', 'B');

    }

    public static void Hanoi(int n, char src, char dst, char tmp) {
        if (n > 0) {
            Hanoi(n, src, tmp, dst);
            System.out.println("Move disk " + n + " from " + src + " to " + dst);
            Hanoi(n, tmp, dst, src);

        } else {
            return;
        }
    }
}
