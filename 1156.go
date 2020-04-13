package leetcode

func maxRepOpt1(text string) int {
	var arr = [26][][2]int{}
	var preChar = text[0]
	var preLeftIdx = 0
	var i int
	for i = 1; i < len(text); i++ {
		if text[i] != preChar {
			arr[preChar-'a'] = append(arr[preChar-'a'], [2]int{preLeftIdx, i-1})
			preLeftIdx = i
		}
		preChar = text[i]
	}
	arr[preChar-'a'] = append(arr[preChar-'a'], [2]int{preLeftIdx, i-1})
	var res int
	for _, intervals := range arr {
		count := len(intervals)
		if count == 0 {
			continue
		}
		var max = intervals[0][1] - intervals[0][0] + 1
		if count > 1 {
			max++
		}
		for i := 1; i < count; i++ {
			if intervals[i][0] - intervals[i-1][1] == 2 {
				t := intervals[i][1] - intervals[i-1][0]
				if count > 2 {
					t++
				}
				if t > max {
					max = t
				}
			} else {
				t := intervals[i][1] - intervals[i][0] + 1
				if count > 1 {
					t++
				}
				if t > max {
					max = t
				}
			}
		}
		if max > res {
			res = max
		}
	}
	return res
}
