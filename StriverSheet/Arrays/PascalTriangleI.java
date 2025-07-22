class Solution {

    public static void main(String[] args) {
        // System.out.println(pascalTriangleI(4, 2));
        int[] r = optimisedPrintRow(5);
        printRow(r);
    }

    public static void printRow(int[] r) {
        for (int i = 0; i < r.length; i++) {
            System.out.println(r[i]);
        }
    }

    public static int pascalTriangleI(int r, int c) {
        r = r - 1;
        c = c - 1;
        return fact(r) / (fact(r - c) * fact(c));
    }

    public static int[] printPascalRow(int r) {
        int[] result = new int[r];
        result[0] = 1;
        for (int c = 1; c < r; c++) {
            result[c] = pascalTriangleI(r, c);
        }
        result[r] = 1;
        return result;
    }

    public static int[] optimisedPrintRow(int r) {
        int[] result = new int[r];
        result[0] = 1;
        result[r - 1] = 1;
        int ans = 1;
        for (int c = 1; c < r; c++) {
            ans = (ans * (r - c)) / c;
            result[c] = ans;
        }
        return result;
    }

    public static int fact(int n) {
        int f = 1;
        // n!/(n-r)!*r!
        for (int i = 1; i <= n; i++) {
            f *= i;
        }
        return f;
    }
}
