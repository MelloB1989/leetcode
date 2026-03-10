package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

func EarliestAppropriateDuration(slots1 []*Interval, slots2 []*Interval, duration int) *Interval {
	result := &Interval{
		Start: -1,
		End:   -1,
	}
	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i].Start < slots1[j].Start
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i].Start < slots2[j].Start
	})
	i, j := 0, 0
	for i < len(slots1) && j < len(slots2) {
		st, en := max(slots1[i].Start, slots2[j].Start), min(slots1[i].End, slots2[j].End)
		if en-st >= duration {
			result.Start = st
			result.End = st + duration
			break
		}
		if slots2[j].End < slots1[i].End {
			j++
		} else {
			i++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	slots1 := []*Interval{
		&Interval{Start: 10, End: 50},
		&Interval{Start: 60, End: 120},
		&Interval{Start: 140, End: 210},
	}
	slots2 := []*Interval{
		&Interval{Start: 0, End: 15},
		&Interval{Start: 60, End: 70},
	}
	fmt.Println(EarliestAppropriateDuration(slots1, slots2, 8))
}
