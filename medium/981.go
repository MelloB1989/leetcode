package main

import (
	"fmt"
)

type TimeMap struct {
	data map[string][]Value
}

type Value struct {
	Timestamp int
	Value     string
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], Value{
		Timestamp: timestamp,
		Value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if pairs, exists := this.data[key]; exists {
		left := 0
		right := len(pairs) - 1

		for left <= right {
			mid := (left + right) / 2
			if pairs[mid].Timestamp == timestamp {
				return pairs[mid].Value
			} else if pairs[mid].Timestamp < timestamp {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if right >= 0 {
			return pairs[right].Value
		}
	}
	return ""
}

func main() {
	// Testing the implementation
	tm := Constructor()

	tm.Set("foo", "bar", 1)
	tm.Set("foo", "baz", 3)

	fmt.Println(tm.Get("foo", 2)) // Should print "bar"
	fmt.Println(tm.Get("foo", 3)) // Should print "baz"
	fmt.Println(tm.Get("foo", 4)) // Should print "baz"
}
