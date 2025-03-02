type TrieNode struct {
    children map[byte]*TrieNode
    isEnd    bool
}

type WordDictionary struct {
    root *TreeNode
}


func Constructor() WordDictionary {
    return WordDictionary{
		root: &TreeNode{
			children: make(map[byte]*TreeNode),
			isEnd: false
		}
	}
}


func (this *WordDictionary) AddWord(word string)  {
	node := this.root
    for i := 0; i < len(word); i++ {
        char := word[i]
        if _, exists := node.children[char]; !exists {
            node.children[char] = &TrieNode{
                children: make(map[byte]*TrieNode),
                isEnd:    false,
            }
        }
        node = node.children[char]
    }
    node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.searchInNode(this.root, word, 0)
}



func (this *WordDictionary) searchInNode(node *TrieNode, word string, index int) bool {
    if index == len(word) {
        return node.isEnd
    }
    
    char := word[index]
    
    // If the current character is '.', we need to check all possible paths
    if char == '.' {
        for _, child := range node.children {
            if this.searchInNode(child, word, index+1) {
                return true
            }
        }
        return false
    }
    
    // Otherwise, just follow the path for the specific character
    nextNode, exists := node.children[char]
    if !exists {
        return false
    }
    
    return this.searchInNode(nextNode, word, index+1)
}