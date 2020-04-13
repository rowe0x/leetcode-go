package leetcode

import "testing"

func Test_isBoomerang(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{points:[][]int{{1,1}, {2,3}, {3, 3}}}, true},
		{"test2", args{points:[][]int{{1,1}, {2,2}, {3, 3}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBoomerang(tt.args.points); got != tt.want {
				t.Errorf("isBoomerang() = %v, want %v", got, tt.want)
			}
		})
	}
}
