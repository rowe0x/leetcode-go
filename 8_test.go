package leetcode

import (
	"testing"
)

func TestMyAoti(t *testing.T) {
	var tables = []struct {
		input string
		expect int
	} {
		{input:"42", expect:42},
		{input:"   -42", expect:-42},
		{input:"word   -42", expect:0},
		{input:"-91283472332", expect:-2147483648},
		{input:"9223372036854775808", expect:2147483647},
	}
	for _, table := range tables {
		actual := myAtoi(table.input)
		if actual != table.expect {
			t.Errorf("expect: %v, found: %v", table.expect, actual)
		}
	}
}
