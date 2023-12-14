package trie

import (
	"strings"
)

// Declaring trie_Node  for creating node in a trie
type trie_Node struct {
	//assigning limit of 26 for child nodes
	childrens [26]*trie_Node
	//declaring a bool variable to check the word end.
	wordEnds bool
}

// Initializing the root of the trie
type Trie struct {
	root *trie_Node
}

// inititlaizing a new trie
func TrieData() *Trie {
	t := new(Trie)
	t.root = new(trie_Node)
	return t
}

// Passing words to trie
func (t *Trie) Insert(word string) {
	current := t.root
	word = strings.Trim(strings.ToLower(word), " ")
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = new(trie_Node)
		}
		current = current.childrens[index]
	}
	current.wordEnds = true
}

// Initializing the search for word in node
func (t *Trie) Search(word string) bool {
	current := t.root
	word = strings.Trim(strings.ToLower(word), " ")
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	return current.wordEnds
}

// Get words that start with prefix
func (t *Trie) GetWordsThatStartWith(prefix string) []string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	words := []string{}
	current := t.root
	prefix = strings.Trim(strings.ToLower(prefix), " ")
	for _, wr := range prefix {
		index := wr - 'a'
		if current.childrens[index] == nil {
			// no words start with this prefix
			return words
		}
		current = current.childrens[index]
	}
	// Not sure if it's uncool to put the type in this function, but it's not used
	// anywhere else so, seems like an ok thing to do
	type StackEntry struct {
		Node   *trie_Node
		Prefix string
	}

	stack := []StackEntry{
		{
			Node:   current,
			Prefix: prefix,
		},
	}
	// The main idea here is to stack the roots and prefixes up and process them in reverse order until
	// There are no more to process. I strongly dislike a never ending for loop, but...
	for {
		if len(stack) == 0 {
			break
		}
		// save working root/prefix for local use and "pop" it off your poor woman's stack
		myCurr := stack[len(stack)-1].Node
		myPre := stack[len(stack)-1].Prefix
		stack = stack[:len(stack)-1]
		if myCurr.wordEnds {
			words = append(words, myPre)
		}
		for _, wr := range alphabet {
			index := wr - 'a'
			if myCurr.childrens[index] == nil {
				continue
			}
			// push this root and associated prefix for later processing
			stack = append(stack, StackEntry{
				Node:   myCurr.childrens[index],
				Prefix: myPre + string(alphabet[index]),
			})
		}
	}
	return words
}

// Do any words start with this prefix
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	prefix = strings.Trim(strings.ToLower(prefix), " ")
	for _, wr := range prefix {
		index := wr - 'a'
		if current.childrens[index] == nil {
			// no words start with this prefix
			return false
		}
		current = current.childrens[index]
	}
	return true
}
