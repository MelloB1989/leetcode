func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	min := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if nums[left] < min {
			min = nums[left]
		}
		if nums[mid] < min {
			min = nums[mid]
		}
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return min
}
