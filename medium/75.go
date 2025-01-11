package main

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
