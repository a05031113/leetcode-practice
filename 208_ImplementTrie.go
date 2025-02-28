type TrieNode struct {
    children map[byte]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    return Trie{
        root: &TrieNode{
            children: make(map[byte]*TrieNode),
            isEnd:    false,
        },
    }
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
    node := this.root
    for i := 0; i < len(word); i++ {
        char := word[i]
        if _, exists := node.children[char]; !exists {
            return false
        }
        node = node.children[char]
    }
    return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
    node := this.root
    for i := 0; i < len(prefix); i++ {
        char := prefix[i]
        if _, exists := node.children[char]; !exists {
            return false
        }
        node = node.children[char]
    }
    return true
}