package leetcode

import "testing"

func TestNumSteps(t *testing.T) {
	var table = []struct {
		input string
		expect int
	}{
		{"1001",7},
	}
	for _, tc := range table {
		actual := numSteps(tc.input)
		if actual != tc.expect {
			t.Errorf("expect %v, found %v", tc.expect, actual)
		}
	}
}

