//Please accept the inputs from the console and print the output
#include <stdio.h>
#include <limits.h>

int maxContiguous(int arr[], int length){
    int max = INT_MIN;
    int sum = 0;
    for(int i=0;i<length;i++){
        sum += arr[i];
        if (sum>max){
            max = sum;
        }
        if (sum < 0){
            sum = 0;
        }
    }
    return max;
}

int main()
{
	// int size;
	// scanf("%d",&size);
	// int arr[size];
	// for(int i=0; i<size; i++){
	//     scanf("%d",&arr[i]);
	// }
	int size =8;
	int arr[] = {10, -3, -4, 7, 6, 5, -4, -1};
	int total_sum;
	int s = maxContiguous(arr, size);

	for(int i=0;i<size;i++){
	    total_sum += arr[i];
        arr[i] = -arr[i];
	}
	total_sum = total_sum + maxContiguous(arr,size);
	if(total_sum>s){
        printf("%d",total_sum);
    } else{
        printf("%d",s);
    }
}
