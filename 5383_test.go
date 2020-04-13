package leetcode

import "testing"

func Test_numOfWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "test1", args: args{1}, want: 12},
		{name: "test2", args: args{2}, want: 54},
		{name: "test3", args: args{5}, want: 30228214},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfWays(tt.args.n); got != tt.want {
				t.Errorf("numOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
