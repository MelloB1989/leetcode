package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: if a person could attend all meetings
 */
func CanAttendMeetings(intervals []*Interval) bool {
	if len(intervals) == 0 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	e := intervals[0].End
	for i, interval := range intervals {
		if i == 0 {
			continue
		}
		if interval.Start < e {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanAttendMeetings([]*Interval{
		&Interval{Start: 35, End: 40},
		&Interval{Start: 1, End: 30},
	}))
}
