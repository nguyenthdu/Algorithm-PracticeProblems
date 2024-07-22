package optimal;

public class optimalIF {
    public static void main(String[] args) {
        // condition("a", "a", "a");
        // conditionOptimal("a", "a", "a");
        conditionOptimal2("a", "a", "a");
    }

    public static void printA() {
        System.out.println("Hello A");
    }

    public static void printB() {
        System.out.println("B");
    }

    public static void condition(String n, String m, String z) {
        if (!n.equals("")) {
            if (!m.equals("")) {
                if (!z.equals("")) {
                    printA();
                } else {
                    System.out.println("z is empty");
                }
            } else {
                System.out.println("m is empty");
            }
        } else {
            System.out.println("n is empty");
        }

    }

    // optimal condition
    public static void conditionOptimal(String n, String m, String z) {
        if (n.equals("") || m.equals("") || z.equals("")) {
            if (n.equals("")) {
                System.out.println("n is empty");
            } else if (m.equals("")) {
                System.out.println("m is empty");
            } else {
                System.out.println("z is empty");
            }
        } else {
            printA();
        }
    }

    // optimal condition
    public static void conditionOptimal2(String n, String m, String z) {
        if (n.equals("")) {
            System.out.println("n is empty");
            return;
        }
        if (m.equals("")) {
            System.out.println("m is empty");
            return;
        }
        if (z.equals("")) {
            System.out.println("z is empty");
            return;
        }
        printA();

    }

    public static void conditionOptimal3(String n, String m, String z) {

        enum Check {
            N, M, Z
        }

    }
}
