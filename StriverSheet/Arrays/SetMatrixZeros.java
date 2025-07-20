class Solution {

    public static void main(String[] args) {
        int[][] matrix = new int[][] {
            { 0, 1, 2, 0 },
            { 3, 4, 5, 2 },
            { 1, 3, 0, 5 },
        }; //{ { 1, 1, 1 }, { 1, 0, 1 }, { 1, 1, 1 } };
        setZeroes(matrix);
    }

    public static void printMatrix(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        for (int row = 0; row < m; row++) {
            for (int col = 0; col < n; col++) {
                System.out.print(matrix[row][col] + ", ");
            }
            System.out.println();
        }
    }

    public static void setZeroes(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        int[] rowMarked = new int[m];
        int[] colMarked = new int[n];
        for (int row = 0; row < m; row++) {
            for (int col = 0; col < n; col++) {
                if (matrix[row][col] == 0) {
                    rowMarked[row] = 1;
                    colMarked[col] = 1;
                }
            }
        }
        // Zero out rows
        for (int i = 0; i < m; i++) {
            if (rowMarked[i] == 1) {
                for (int j = 0; j < n; j++) matrix[i][j] = 0;
            }
        }
        // Zero out columns
        for (int j = 0; j < n; j++) {
            if (colMarked[j] == 1) {
                for (int i = 0; i < m; i++) matrix[i][j] = 0;
            }
        }
        printMatrix(matrix);
    }
}
