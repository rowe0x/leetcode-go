package leetcode

import (
	"testing"
)

func Test_findLongestChain(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name:"test1", args:args{pairs:[][]int{{-10,-8}, {8,9}, {-5,0},{6,10},{-6,-4},{1,7},{9,10},{-4,7}}}, want:4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestChain(tt.args.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}