package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	faces := []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
	time := ""

	var findTimes func(time string, limit int, i int, hours int, mins int)
	findTimes = func(time string, limit int, i int, hours int, mins int) {
		if limit == 0 {
			if hours > 11 || mins > 59 {
				return
			}
			if mins < 10 {
				time = fmt.Sprintf("%d:0%d", hours, mins)
			} else {
				time = fmt.Sprintf("%d:%d", hours, mins)
			}
			result = append(result, time)
		}

		for j := i; j < len(faces); j++ {
			if j < 4 {
				hours += faces[j]
			} else {
				mins += faces[j]
			}
			findTimes(time, limit-1, (j + 1), hours, mins)
			if j < 4 {
				hours -= faces[j]
			} else {
				mins -= faces[j]
			}
		}
	}

	findTimes(time, turnedOn, 0, 0, 0)

	return result
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
