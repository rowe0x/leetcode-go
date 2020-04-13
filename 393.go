package leetcode

func validUtf8(data []int) bool {
	nums := len(data)
	if nums == 0 {
		return true
	}
	if nums == 1 {
		return data[0]&0x80 == 0
	}
	fb := data[0]
	if fb&0x80 == 0 {
		return validUtf8(data[1:])
	}
	if (fb<<1)&0x80 == 0 {
		return false
	}
	tBit := 0
	if (fb<<2)&0x80 != 0 {
		tBit = 1
	}
	fBit := 0
	if (fb<<3)&0x80 != 0 {
		fBit = 1
	}
	fiveBit := 0
	if (fb<<4)&0x80 != 0 {
		fiveBit = 1
	}

	sb := data[1]
	if sb&0x80 == 0 {
		return false
	}
	if (sb<<1)&0x80 != 0 {
		return false
	}
	if tBit == 0 { // 2 bytes
		return validUtf8(data[2:])
	}

	if nums < 3 {
		return false
	}
	tb := data[2]
	if tb&0x80 == 0 {
		return false
	}
	if (tb<<1)&0x80 != 0 {
		return false
	}
	if fBit == 0 { // 3 bytes
		return validUtf8(data[3:])
	}

	if nums < 4 {
		return false
	}
	fourb := data[3]
	if fourb&0x80 == 0 {
		return false
	}
	if (fourb<<1)&0x80 != 0 {
		return false
	}
	if fiveBit == 1 {
		return false
	}
	return validUtf8(data[4:])
}
