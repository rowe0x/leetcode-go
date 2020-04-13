package leetcode

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	var ans []string
	sort.Strings(folder)
	ans = append(ans, folder[0])
	prefix := folder[0] + "/"
	for i := 1; i < len(folder); i++ {
		if !strings.HasPrefix(folder[i]+"/", prefix) {
			ans = append(ans, folder[i])
			prefix = folder[i] + "/"
		}
	}
	return ans
}
