package leetcode

import (
	"time"
)

func dayOfYear(date string) int {
	t0, _ := time.Parse("2006", date[0:4])
	t, _ := time.Parse("2006-01-02", date)
	return int(t.Sub(t0).Hours()/24)+1
}
