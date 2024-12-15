package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 > n2 {
		nums1, nums2 = nums2, nums1
		n1, n2 = n2, n1
	}
	left, right := 0, n1
	for left <= right {
		i := (left + right) / 2
		j := (n1+n2+1)/2 - i

		maxLeft1 := math.MinInt64
		if i > 0 {
			maxLeft1 = nums1[i-1]
		}

		minRight1 := math.MaxInt64
		if i < n1 {
			minRight1 = nums1[i]
		}

		maxLeft2 := math.MinInt64
		if j > 0 {
			maxLeft2 = nums2[j-1]
		}

		minRight2 := math.MaxInt64
		if j < n2 {
			minRight2 = nums2[j]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (n1+n2)%2 == 0 {
				return (float64(math.Max(float64(maxLeft1), float64(maxLeft2))) + float64(math.Min(float64(minRight1), float64(minRight2)))) / 2
			} else {
				return math.Max(float64(maxLeft1), float64(maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return float64(0)
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
