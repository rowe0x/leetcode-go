package leetcode

import (
	"testing"
)

func Test_maxRepOpt1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"ababa"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepOpt1(tt.args.text); got != tt.want {
				t.Errorf("maxRepOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}