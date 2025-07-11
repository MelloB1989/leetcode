/*
Problem:

You are give a array of integers where each element shows the magnitude of the water current of a river. If the current is negative then the water current direction is left and if the current is positive then the direction is right. If the current is 0 then the water is stagnant.
Modify and return the array by replacing all the stagnant element with the stronger current adjacent.

Example:
array: {2,0,0,-4,0,8}
output: {2,-4,-4,8,8}
*/

class RiverCurrent {

    int[] currents = new int[] { 2, 0, 0, -4, 0, 8 };
    int len = 6;

    public int abs(int x) {
        if (x < 0) {
            return -x;
        }
        return x;
    }

    public void modify() {
        for (int i = 0; i < len; i++) {
            if (currents[i] == 0) {
                int rightCurrent = i;
                int leftCurrent = i;

                while (currents[leftCurrent] == 0) {
                    leftCurrent++;
                    if (leftCurrent >= len) {
                        leftCurrent = len - 1;
                    }
                }

                while (currents[rightCurrent] == 0) {
                    rightCurrent--;
                    if (rightCurrent <= 0) {
                        rightCurrent = 0;
                    }
                }

                int strongest = abs(currents[leftCurrent]) >
                    abs(currents[rightCurrent])
                    ? currents[leftCurrent]
                    : currents[rightCurrent];

                while (currents[i] == 0) {
                    currents[i] = strongest;
                    i++;
                }
            }
        }
    }
}
