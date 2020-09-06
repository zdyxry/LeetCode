package main

type TimeMap struct {
	timestamp map[string][]int
	realValue map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		timestamp: make(map[string][]int),
		realValue: make(map[string][]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timestamp[key] = append(this.timestamp[key], timestamp)
	this.realValue[key] = append(this.realValue[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.timestamp[key]) == 0 {
		return ""
	}

	index := findFirstLEIndex(this.timestamp[key], timestamp)
	if index < 0 {
		return ""
	}
	return this.realValue[key][index]
}

func findFirstLEIndex(nums []int, n int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] > n {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}
