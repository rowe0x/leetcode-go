package leetcode

import (
	"strings"
)

func printVertically(s string) []string {
	var ans []string
	var iterCount int
	sp := strings.Split(s, " ")
	for _, t := range sp {
		if iterCount < len(t) {
			iterCount = len(t)
		}
	}
	for i := 0; i < iterCount; i++ {
		var c []byte
		for _, t := range sp {
			if i < len(t) {
				c = append(c, t[i])
			} else {
				c = append(c, ' ')
			}
		}
		ans = append(ans, strings.TrimRight(string(c), " "))
	}
	return ans
}
