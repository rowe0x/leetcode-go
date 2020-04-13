package leetcode

func simplifyPath(path string) string {
	var slashes []int
	var ans []byte
	start, end := 0, len(path)-1
	var dotCount int
	for start <= end {
		switch path[start] {
		case '/':
			if dotCount == 1 {
				ans = ans[0:len(ans)-1]
			} else if dotCount == 2 {
				if len(slashes) > 1 {
					ans = ans[0:slashes[len(slashes)-2]]
					slashes = slashes[0:len(slashes)-2]
				} else {
					ans = ans[0:0]
					slashes = slashes[0:0]
				}
			}
			if len(ans) == 0 || ans[len(ans)-1] != '/' {
				ans = append(ans, '/')
				slashes = append(slashes, len(ans)-1)
			}
			dotCount = 0
			start++
		case '.':
			ans = append(ans, '.')
			dotCount++
			start++
		default:
			ans = append(ans, path[start])
			dotCount = 0
			start++
		}
	}
	if dotCount == 1 {
		ans = ans[0:len(ans)-1]
	} else if dotCount == 2 {
		if len(slashes) > 1 {
			ans = ans[0:slashes[len(slashes)-2]]
			ans = append(ans, '/')
		} else {
			ans = ans[0:0]
			ans = append(ans, '/')
		}
	}
	if len(ans) > 1 && ans[len(ans)-1] == '/' {
		ans = ans[:len(ans)-1]
	}
	return string(ans)
}
