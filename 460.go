package leetcode

import "fmt"

type LFUCache struct {
	data map[int]*Node
	cap int
	head *Node
}

type Node struct {
	Key int
	Val int
	Freq int
	Next *Node
}

func NewLFUCaChe(capacity int) LFUCache {
	return LFUCache{data: make(map[int]*Node), cap: capacity}
}


func (c *LFUCache) Get(key int) int {
	if c.cap == 0 {
		return -1
	}
	fmt.Printf("%+v\n", c.data)
	node, ok := c.data[key]
	if !ok {
		return -1
	}

	if node.Next == nil {
		node.Freq++
		return node.Val
	}
	// delete node
	next := node.Next
	val, freq := node.Val, node.Freq+1

	node.Key = node.Next.Key
	node.Val = node.Next.Val
	node.Freq = node.Next.Freq
	node.Next = node.Next.Next

	next.Key, next.Val, next.Freq = key, val, freq
	update(node, next)
	c.data[key] = next
	c.data[node.Key] = node
	return val
}

func update(node, next *Node) {
	for node.Next != nil && node.Freq <= next.Freq {
		node = node.Next
	}
	if node.Freq <= next.Freq {
		node.Next = next
		next.Next = nil
	} else {
		next.Next = node.Next
		node.Next = next
	}
}

func (c *LFUCache) Put(key int, value int)  {
	if c.cap == 0 {
		return
	}
	node, ok := c.data[key]
	if !ok {
		if c.cap == len(c.data) {
			hk := c.head.Key
			c.head = c.head.Next
			delete(c.data, hk)
		}
		if c.head == nil {
			c.head = &Node{Val: value, Freq: 1, Key: key}
			c.data[key] = c.head
			return
		}
		if c.head.Freq > 1 {
			c.head = &Node{
				Key:  key,
				Val:  value,
				Freq: 1,
				Next: c.head,
			}
			c.data[key] = c.head
			return
		}

		node := c.head
		next := &Node{Val: value, Freq: 1, Key: key}
		update(node, next)
		c.data[key] = next
		return
	}
	if node.Next == nil {
		node.Freq++
		node.Val = value
		return
	}
	// delete node
	next := node.Next
	val, freq := value, node.Freq+1

	node.Key = node.Next.Key
	node.Val = node.Next.Val
	node.Freq = node.Next.Freq
	node.Next = node.Next.Next

	next.Key, next.Val, node.Freq = key, val, freq
	update(node, next)
	c.data[key] = next
	c.data[node.Key] = node
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
