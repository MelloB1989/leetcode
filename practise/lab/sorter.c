#include <stdio.h>

void sorter(int arr[], int size){
    for (int i=0;i<size;i++){
        for (int j=i; j<size; j++){
            if (arr[i]>arr[j]){
                int t = arr[j];
                arr[j] = arr[i];
                arr[i] = t;
            }
        }
    }
    for (int i=0; i<size; i++){
        printf("%d ",arr[i]);
    }
}

int main(){
    int size = 5;
    int arr[] = {0,2,1,2,0};
    sorter(arr,size);
    return 0;
}

// int main(){
//     int size, k;
//     scanf("%d", &size);
//     int arr[size];
//     for (int i=0; i<size; i++) scanf("%d", &arr[i]);
//     sorter(arr,size);
// }
