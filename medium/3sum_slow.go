package main

import "fmt"

func threeSum(nums []int) [][]int {
	var r [][]int
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		for j < length-1 {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			k := length - 1
			for k > j {
				if k < length-1 && nums[k] == nums[k+1] {
					k--
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					r = append(r, []int{nums[i], nums[j], nums[k]})
					break
				}
				k--
			}
			j++
		}
	}
	return r
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
