package tree

// TrieNode represents a node in the trie
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// NewTrie creates a new TrieNode
func NewTrie() *TrieNode {
	return &TrieNode{}
}

// Insert a word into the trie
func (t *TrieNode) Insert(word string) {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &TrieNode{}
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

// Search a word in the trie
func (t *TrieNode) Search(word string) bool {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return node.isEnd
}

// Returns if there is any word in the trie that starts with the given prefix
func (t *TrieNode) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return true
}
