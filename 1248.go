package leetcode

func numberOfSubarrays(nums []int, k int) int {
	windows := make([]int, 1)
	var lastIndex int
	for _, n := range nums {
		if n % 2 == 0 {
			windows[lastIndex]++
		} else {
			windows = append(windows, -1, 0)
			lastIndex += 2
		}
	}
	left, right := 0, 0
	found := false
	for left = 0; left < len(windows); left++ {
		if nums[left] == -1 {
			found = true
			break
		}
	}
	if !found {
		return 0
	}
	var ans int
	if k == 1 {
		for i := left; i < len(windows); left++ {
			if nums[i] == -1 {
				l := 1
				if left-1 >= 0 {
					l = windows[left-1]
				}
				r := 1
				if right+1 < len(windows) {
					r = windows[right+1]
				}
				ans += l * r
			}
		}
		return ans
	}
	oddNums := 1
	for right = left+1; right < len(windows); right++ {
		if windows[right] == -1 {
			oddNums++
			if oddNums == k {
				l := 1
				if left-1 >= 0 {
					l = windows[left-1]+1
				}
				r := 1
				if right+1 < len(windows) {
					r = windows[right+1]
				}
				ans += l * r
				left = right
				oddNums--
			}
		}
	}
	return ans
}
