class Solution {
    // public static String mergeAlternately(String word1, String word2) {
    // String result = "";
    // int length = word1.length() + word2.length();
    // // if (word1.length() > word2.length())
    // // length = word1.length();
    // // else
    // // length = word2.length();
    // int k = 0;
    // int l = 0;
    // for (int i = 0; i <= length; i++) {
    // if (i % 2 == 0) {
    // if (l < word1.length()) {
    // result += word1.charAt(l);
    // l++;
    // }
    // } else {
    // if (k < word2.length()) {
    // result += word2.charAt(k);
    // k++;
    // }

    // }

    // }
    // result += word1.substring(l);
    // result += word2.substring(k);

    // return result;

    // }
    public static String mergeAlternately(String word1, String word2) {
        StringBuilder result = new StringBuilder();
        int i = 0;
        while (i < word1.length() || i < word2.length()) {
            if (i < word1.length())
                result.append(word1.charAt(i));// a
            if (i < word2.length())
                result.append(word2.charAt(i));// d
            i++;// ab be c
            // them tung cap word1[0] word2[0]
        }
        return result.toString();

    }

    public static void main(String[] args) {
        String a = "abc";
        String b = "de";
        System.out.println(mergeAlternately(a, b));
    }
}
