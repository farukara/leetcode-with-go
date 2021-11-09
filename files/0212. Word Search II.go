package main
//leetcode 212. Word Search II

import (
    "fmt"
)
const ALPHABET_SIZE = 26

type trieNode struct {
    childrens [ALPHABET_SIZE]*trieNode
    isWordEnd bool
}

type trie struct {
    root *trieNode
}

func initTrie() *trie {
    return &trie{
        root: &trieNode{},
    }
}

func (t *trie) insert(word string) {
    wordLength := len(word)
    current := t.root
    for i:=0; i<wordLength; i++ {
        index := word[i]-'a'
        if current.childrens[index] == nil {
            current.childrens[index] = &trieNode{}
        }
        current = current.childrens[index]
    }
    current.isWordEnd = true
}

func (t *trie) find(word string) bool {
    wordLength := len(word)
    current := t.root
    for i:=0; i<wordLength; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            return false
        }
        current = current.childrens[index]
    }
    return current.isWordEnd
}

func findWords(board [][]string, words []string) []string {
    trie := initTrie()
    for i:=0; i<len(words); i++ {
        trie.insert(words[i])
    }
    current := trie.root
    for row:=0; row<len(board); row++ {
        for column:=0; column<len(board[0]); column++ {
            index := int([]byte(board[row][column])[0])-'a' //TODO: find a more elegant way to find index
            if current.childrens[index] != nil {
                //FIX: continue from here. find a way to recurcively find neighbors and continue

                fmt.Println(board[row][column])
            }
        }
    }
    result := []string{"q"}
    return result
}

func main() {
    board := [][]string{{"o","a","a","n"}, {"e","t","a","e"}, {"i","h","k","r"}, {"i","f","l","v"}}
    words := []string{"oath","pea","eat","rain"}
    fmt.Println(findWords(board, words))
}
