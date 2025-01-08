package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	if strNums[0] == "0" {
		return "0"
	}

	result := ""
	for _, str := range strNums {
		result += str
	}

	return result
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
