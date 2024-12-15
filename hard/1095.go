package main

import (
	"fmt"
)

type MountainArray struct {
	Arr []int
}

func (this *MountainArray) get(index int) int {
	return this.Arr[index]
}
func (this *MountainArray) length() int {
	return len(this.Arr)
}
func Constructor() *MountainArray {
	return &MountainArray{
		Arr: []int{1, 2, 3, 4, 5, 3, 1},
		// Arr: []int{1, 5, 2},
	}
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	search := func(left, right, target int, inc bool) int {
		for left <= right {
			mid := (left + right) / 2
			m := mountainArr.get(mid)
			if m == target {
				return mid
			}
			if inc {
				if m < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else {
				if m < target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
		return -1
	}
	left, right := 0, n-1
	peak := 0
	for left < right {
		mid := (left + right) / 2
		fmt.Println(left, mid, right)
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	peak = left
	//Search left side
	l := search(0, peak, target, true)
	if l != -1 {
		return l
	}
	//Search right side
	return search(peak, n-1, target, false)
}

func main() {
	ma := Constructor()
	fmt.Println(findInMountainArray(3, ma))
}
