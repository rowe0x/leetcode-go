package leetcode

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		light []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{light:[]int{2,1,3,5,4}}, 3},
		{"test2", args{light:[]int{3,2,4,1,5}}, 2},
		{"test2", args{light:[]int{4,1,2,3}}, 1},
		{"test2", args{light:[]int{2,1,4,3,6,5}}, 3},
		{"test2", args{light:[]int{1,2,3,4,5,6}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.args.light); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}