package leetcode

func romanToInt(s string) int {
	var stack =  [1]int{}
	var ans int
	for i := 0; i < len(s); i++ {
		switch i {
		case 'I':
			ans += 1
			stack[0] = 1
		case 'V':
			if stack[0] == 1 {
				ans--
				ans += 4
			} else {
				ans += 5
			}
		case 'X':
			if stack[0] == 1 {
				ans--
				ans += 9
			} else {
				ans += 10
			}
			stack[0] = 10
		case 'L':
			if stack[0] == 10 {
				ans -= 10
				ans += 40
			} else {
				ans += 50
			}
		case 'C':
			if stack[0] == 10 {
				ans -= 10
				ans += 90
			} else {
				ans += 100
			}
			stack[0] = 100
		case 'D':
			if stack[0] == 100 {
				ans -= 100
				ans += 400
			} else {
				ans += 500
			}
		case 'M':
			if stack[0] == 100 {
				ans -= 100
				ans += 900
			} else {
				ans += 1000
			}
		}
	}
	return ans
}
