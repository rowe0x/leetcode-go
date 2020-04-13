package leetcode

import "testing"

func TestLFUCache_Get(t *testing.T) {
	var val int
	c := NewLFUCaChe(2)
	c.Put(2,1)
	c.Put(1,1)
	c.Put(2,3)
	c.Put(4,1)
	var table = [][]int {
		{1,-1},
		{2,3},
	}
	for _, item := range table {
		val = c.Get(item[0])
		if val != item[1] {
			t.Errorf("key %d expect %d got %d", item[0], item[1], val)
		}
	}
}
