package leetcode

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	var tables = []struct {
		input string
		expect string
	}{
		// {input:"/home//foo/", expect:"/home/foo"},
		//{input:"/a/./b/../../c/", expect:"/c"},
		//{input:"/a/../../b/../c//.//", expect:"/c"},
		//{input:"/a//b////c/d//././/..", expect:"/a/b/c"},
		// {input:"/..", expect:"/"},
		{input:"///eHx/..", expect:"/"},
	}
	for _, table := range tables {
		actual := simplifyPath(table.input)
		if actual != table.expect {
			t.Errorf("input: %v, expect: %v, found: %v", table.input, table.expect, actual)
		}
	}
}
