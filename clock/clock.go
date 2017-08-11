package clock

import "fmt"

const testVersion = 4

type Clock struct {
	h, m int
}

func rolloverMinutes(min int) (int, int) {
	var hours int
	if min < 0 {
		for min < 0 {
			hours -= 1
			min += 60
		}
	}
	if min >= 60 {
		for min >= 60 {
			hours += 1
			min -= 60
		}
	}
	return hours, min
}

func rolloverHours(hrs int) int {
	if hrs >= 24 {
		for hrs >= 24 {
			hrs -= 24
		}
	}
	if hrs < 0 {
		for hrs < 0 {
			hrs += 24
		}
	}
	return hrs
}

func New(hrs, min int) Clock {
	h, m := rolloverMinutes(min)
	hrs += h
	h = rolloverHours(hrs)
	clock := Clock{
		h,
		m,
	}
	return clock
}

func (c Clock) String() string {
	hr, min := rolloverMinutes(c.m)
	hrs := hr + c.h
	hrs = rolloverHours(hrs)
	var hoursString string
	var minString string
	if hrs < 10 {
		hoursString = fmt.Sprintf("0%v", hrs)
	} else {
		hoursString = fmt.Sprintf("%v", hrs)
	}
	if min < 10 {
		minString = fmt.Sprintf("0%v", min)
	} else {
		minString = fmt.Sprintf("%v", min)
	}
	return fmt.Sprintf("%s:%s", hoursString, minString)
}

func (c Clock) Add(minutes int) Clock {
	hrsToAdd, minToAdd := rolloverMinutes(minutes)
	if hrsToAdd >= 24 || hrsToAdd < 0 {
		hrsToAdd = rolloverHours(hrsToAdd)
	}
	c.h += hrsToAdd
	c.m += minToAdd
	return c
}
