package leetcode

type DinnerPlates struct {
	stacks    [][]int
	capacity  int
	rightMost int
	leftMost  int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{stacks: make([][]int, 2), capacity: capacity, rightMost: 0, leftMost: 0}
}

func (this *DinnerPlates) Push(val int) {
	var idx int
	for idx = this.leftMost; ; idx++ {
		if len(this.stacks)-1 < idx {
			this.stacks = append(this.stacks, []int{})
		}
		// fmt.Println(this.stacks,idx)
		if len(this.stacks[idx]) == this.capacity {
			continue
		}
		this.stacks[idx] = append(this.stacks[idx], val)
		break
	}
	this.leftMost = idx
	if this.leftMost > this.rightMost {
		this.rightMost = this.leftMost
	}
}

func (this *DinnerPlates) Pop() int {
	if len(this.stacks) == 0 {
		return -1
	}
	var idx int
	var value int
	for idx = this.rightMost; idx >= 0; idx-- {
		if len(this.stacks[idx]) == 0 {
			continue
		}
		value = this.stacks[idx][len(this.stacks[idx])-1]
		this.stacks[idx] = this.stacks[idx][0 : len(this.stacks[idx])-1]
		this.rightMost = idx
		if this.rightMost < this.leftMost {
			this.leftMost = this.rightMost
		}
		return value
	}
	this.rightMost = 0
	this.leftMost = 0
	return -1
}

func (this *DinnerPlates) PopAtStack(index int) int {

	if index > this.rightMost {
		return -1
	}
	if len(this.stacks[index]) == 0 {
		return -1
	}
	value := this.stacks[index][len(this.stacks[index])-1]
	this.stacks[index] = this.stacks[index][0 : len(this.stacks[index])-1]
	return value
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
