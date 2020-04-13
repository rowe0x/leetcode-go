package leetcode

type TrieNode struct {
	Char byte
	Children map[byte]*TrieNode
}

func addWord(root *TrieNode, word string) {
	node := root
	for i := 0; i < len(word); i++ {
		if _, ok := node.Children[word[i]]; !ok {
			newNode := &TrieNode{word[i], make(map[byte]*TrieNode)}
			node.Children[word[i]] = newNode
			node = newNode
		}
	}
}

func count(root *TrieNode) int {
	if root == nil {
		return 1
	}
	var ans int
	for _, child := range root.Children {
		ans += count(child)
	}
	return ans
}

func minimumLengthEncoding(words []string) int {
	root := &TrieNode{Children: map[byte]*TrieNode{}}
	for _, word := range words {
		addWord(root, reverse(word))
	}
	return count(root)
}

func reverse(word string) string {
	ans := make([]byte, len(word))
	for i := len(word)-1; i >= 0; i-- {
		ans[len(word)-1-i] = word[i]
	}
	return string(ans)
}
