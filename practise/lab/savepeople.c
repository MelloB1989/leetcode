#include <stdio.h>

int save(int arr[], int size, int limit) {
    for (int i=0;i<size;i++){
        for (int j=i; j<size; j++){
            if (arr[i]>arr[j]){
                int t = arr[j];
                arr[j] = arr[i];
                arr[i] = t;
            }
        }
    }

    int i = 0;
    int j = size - 1;
    int boats = 0;

    while (i <= j) {
        if (arr[i] + arr[j] <= limit) {
            i++;
            j--;
        } else {
            j--;
        }
        boats++;
    }

    return boats;
}

int main(){
    int size = 10;
    int arr[] = {523, 480, 865, 552, 865, 354, 480, 744, 27, 744};
    int limit = 999;
    printf("\n%d",save(arr,size, limit));
    return 0;
}

// int main(){
//     int size, limit;
//     scanf("%d", &size);
//     int arr[size];
//     scanf("%d", &limit);
//     for (int i=0; i<size; i++) scanf("%d", &arr[i]);
//     printf("%d",save(arr,size, limit));
// }
