package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

func MinMeetingRooms(intervals []*Interval) int {
	roomReq := 0
	if len(intervals) == 0 {
		return roomReq
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	en := intervals[0]
	for _, interval := range intervals {
		if interval.Start <
	}
	return roomReq
}

func main() {
	fmt.Println(MinMeetingRooms([]*Interval{
		&Interval{Start: 0, End: 30},
		&Interval{Start: 5, End: 10},
		&Interval{Start: 15, End: 20},
	}))
}
