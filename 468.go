package leetcode

import "strings"

func validIPAddress(IP string) string {
	if ipv4(IP) {
		return "IPv4"
	}
	if ipv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func ipv4(ip string) bool {
	l := strings.Split(ip, ".")
	if len(l) != 4 {
		return false
	}
	for i := 0; i < 4; i++ {
		if len(l[i]) == 0 {
			return false
		}
		var leadZero bool
		var num int
		for j := 0; j < len(l[i]); j++ {
			switch l[i][j] {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if leadZero {
					return false
				}
				num = 10*num + int(l[i][j]-'0')
			case '0':
				if j == 0 {
					leadZero = true
				} else if leadZero {
					return false
				}
			default:
				return false
			}
		}
		if num > 255 {
			return false
		}
	}
	return true
}

func ipv6(ip string) bool {
	l := strings.Split(ip, ":")
	if len(l) != 8 {
		return false
	}
	for i := 0; i < 8; i++ {
		if len(l[i]) > 4 {
			return false
		}
		if len(l[i]) == 0 {
			return false
		}
		for j := 0; j < len(l[i]); j++ {
			switch l[i][j] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			case 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
				continue
			default:
				return false
			}
		}
	}
	return true
}
