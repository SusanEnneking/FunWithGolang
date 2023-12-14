package trie

import (
	"fmt"
	"slices"
	"testing"
)

func TestTrieWordsThatStartWith(t *testing.T) {
	wordList := []string{"apple", "ant", "ape", "cat", "car", "cart", "curtsy", "dog", "dainty", "dart"}
	testTrie := TrieData()
	for _, word := range wordList {
		testTrie.Insert(word)
	}
	prefix := "ap"
	matchingWords := testTrie.GetWordsThatStartWith(prefix)
	if len(matchingWords) != 2 {
		t.Error(fmt.Printf("Expected 2 for prefix: %s, but got %d.", prefix, len(matchingWords)))
	}
	checkWordsAreInOriginalList(prefix, matchingWords, wordList, t)

	prefix = "an"
	matchingWords = testTrie.GetWordsThatStartWith(prefix)
	if len(matchingWords) != 1 {
		t.Error(fmt.Printf("Expected 1 for prefix: %s, but got %d.", prefix, len(matchingWords)))
	}
	checkWordsAreInOriginalList(prefix, matchingWords, wordList, t)

	prefix = "cu"
	matchingWords = testTrie.GetWordsThatStartWith(prefix)
	if len(matchingWords) != 1 {
		t.Error(fmt.Printf("Expected 0 for prefix: %s, but got %d.", prefix, len(matchingWords)))
	}
	checkWordsAreInOriginalList(prefix, matchingWords, wordList, t)

	prefix = "d"
	matchingWords = testTrie.GetWordsThatStartWith(prefix)
	if len(matchingWords) != 3 {
		t.Error(fmt.Printf("Expected 3 for prefix: %s, but got %d.", prefix, len(matchingWords)))
	}
	checkWordsAreInOriginalList(prefix, matchingWords, wordList, t)

	prefix = "f"
	matchingWords = testTrie.GetWordsThatStartWith(prefix)
	if len(matchingWords) != 0 {
		t.Error(fmt.Printf("Expected 0 for prefix: %s, but got %d.", prefix, len(matchingWords)))
	}
	checkWordsAreInOriginalList(prefix, matchingWords, wordList, t)
	prefix = ""
	matchingWords = testTrie.GetWordsThatStartWith(prefix)
	if len(matchingWords) != len(wordList) {
		t.Error(fmt.Printf("Expected 0 for prefix: %s, but got %d.", prefix, len(matchingWords)))
	}
	checkWordsAreInOriginalList(prefix, matchingWords, wordList, t)
}

func checkWordsAreInOriginalList(prefix string, words []string, originalList []string, t *testing.T) {
	for _, word := range words {
		if !slices.Contains(originalList, word) {
			t.Error(fmt.Printf("Some returned words are not in original List for prefix: %s.", prefix))
		}
	}
}

func TestTrieStartsWith(t *testing.T) {
	wordList := []string{"apple", "ant", "ape", "cat", "car", "cart", "curtsy", "dog", "dainty", "dart"}
	testTrie := TrieData()
	for _, word := range wordList {
		testTrie.Insert(word)
	}
	prefix := "ap"
	matchingWords := testTrie.StartsWith(prefix)
	if !matchingWords {
		t.Error(fmt.Printf("Expected true for prefix: %s, but got %t.", prefix, matchingWords))
	}
	prefix = "apa"
	matchingWords = testTrie.StartsWith(prefix)
	if matchingWords {
		t.Error(fmt.Printf("Expected false for prefix: %s, but got %t.", prefix, matchingWords))
	}
}

func TestSearch(t *testing.T) {
	wordList := []string{"apple", "ant", "ape", "cat", "car", "cart", "curtsy", "dog", "dainty", "dart"}
	testTrie := TrieData()
	for _, word := range wordList {
		testTrie.Insert(word)
	}
	word := "curtsy"
	matchingWords := testTrie.Search(word)
	if !matchingWords {
		t.Error(fmt.Printf("Expected true for: %s, but got %t.", word, matchingWords))
	}
	word = "elephant"
	matchingWords = testTrie.Search(word)
	if matchingWords {
		t.Error(fmt.Printf("Expected false for prefix: %s, but got %t.", word, matchingWords))
	}

}
