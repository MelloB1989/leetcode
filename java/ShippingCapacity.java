import java.io.*;
import java.util.*;

public class ShippingCapacity {

    public static boolean canShip(int[] weights, int days, int capacity) {
        int passedDays = 1;
        int currentLoad = 0;

        for (int i = 0; i < weights.length; i++) {
            if (currentLoad + weights[i] > capacity) {
                passedDays += 1;
                currentLoad = 0;
            }
            currentLoad += weights[i];

            if (passedDays > days) return false;
        }

        return true;
    }

    public static int findMinCap(int[] weights, int days) {
        int maxWeight = weights[0];
        int sumWeights = 0;
        for (int i = 0; i < weights.length; i++) {
            if (maxWeight < weights[i]) maxWeight = weights[i];
            sumWeights += weights[i];
        }

        int l = maxWeight;
        int r = sumWeights;

        while (l < r) {
            int mid = (l + r) / 2;
            if (canShip(weights, days, mid)) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }

        return l;
    }

    public static void main(String[] args) {
        int[] weights = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        int days = 5;
        System.out.println(findMinCap(weights, days));
    }
}
