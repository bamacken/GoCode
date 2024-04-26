package tree

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

// Constructor
func (t *Trie) Constructor() *Trie {
	return &Trie{root: &TrieNode{}}
}

/*
Define Trie Node Structure:

Define a TrieNode class:
Include a children attribute, which is a dictionary mapping each character to another TrieNode.
Include an isEndOfWord attribute, a boolean that indicates if the node marks the end of a word.
Initialize the Trie:

Create the root node:
Set the root node as a new instance of TrieNode.
Initialize isEndOfWord to false and children to an empty dictionary.
Function to Insert a Word:

Start at the root node.
For each character in the word:
If the character is not in the current node’s children, add a new TrieNode as its value.
Move to the TrieNode associated with the character.
After the last character:
Set the isEndOfWord of the current node to true.
Function to Search for a Word:

Start at the root node.
For each character in the word:
Check if the character exists in the current node’s children.
If it does, move to the TrieNode associated with the character.
If it doesn't, return false because the word is not in the trie.
After processing all characters:
Return the value of isEndOfWord at the current node. This will be true if the whole word is present and false if only a prefix is found.
(Optional) Function to Check for a Prefix:

Start at the root node.
For each character in the prefix:
Check if the character exists in the current node’s children.
If it does, move to the TrieNode associated with the character.
If it doesn't, return false because the prefix is not present.
If all characters in the prefix are found:
Return true indicating the prefix exists in the trie.
*/
