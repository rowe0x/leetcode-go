package leetcode

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	initial, _ := time.Parse("20060102","19710103")
	end, _ := time.Parse("20060102", fmt.Sprintf("%d%02d%02d",year,month, day))
	days := int(end.Sub(initial).Hours() / 24)
	week := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return week[days%7]
}
