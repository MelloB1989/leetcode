#include <stdio.h>

int numberGame(int arr[], int k, int size){
    for (int i=0; i<size; i++){
        if (arr[i] == k){
            k = k*2;
        }
    }
    return k;
}

int main(){
    int size = 5;
    int arr[] = {5,3,6,1,12};
    int k = 3;
    printf("%d",numberGame(arr, k, size));
    return 0;
}

// int main(){
//     int size, k;
//     scanf("%d", &size);
//     int arr[size];
//     for (int i=0; i<size; i++) scanf("%d", &arr[i]);
//     scanf("%d",&k);
//     printf("%d",numberGame(arr, k, size));
// }
