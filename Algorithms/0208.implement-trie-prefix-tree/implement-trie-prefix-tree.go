package problem0208

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Trie 是便于 word 插入与查找的数据结构
//type Trie struct {
//	val  byte
//	sons [26]*Trie
//	end  int
//}
//
//func Constructor() Trie {
//	return Trie{}
//}
//
//func (this *Trie) Insert(word string) {
//	node := this
//	size := len(word)
//	for i := 0; i < size; i++ {
//		idx := word[i] - 'a'
//		if node.sons[idx] == nil {
//			node.sons[idx] = &Trie{val: word[i]}
//		}
//
//		node = node.sons[idx]
//	}
//
//	node.end++
//}
//
//func (this *Trie) Search(word string) bool {
//	node := this
//	size := len(word)
//	for i := 0; i < size; i++ {
//		idx := word[i] - 'a'
//		if node.sons[idx] == nil {
//			return false
//		}
//		node = node.sons[idx]
//	}
//
//	if node.end > 0 {
//		return true
//	}
//
//	return false
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	node := this
//	size := len(prefix)
//	for i := 0; i < size; i++ {
//		idx := prefix[i] - 'a'
//		if node.sons[idx] == nil {
//			return false
//		}
//		node = node.sons[idx]
//	}
//
//	return true
//}


type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor208() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}