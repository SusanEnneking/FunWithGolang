package main

import (
	"encoding/json"
	"net/http"

	"github.com/SusanEnneking/PlayWithGo/server/api/exercises"
	"github.com/SusanEnneking/PlayWithGo/server/trie"
)

type WordInfo struct {
	Words  []string
	Prefix string
}

type ReturnValues struct {
	Errcode int         `json:"errcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/diagonal", diagonalHandler)
	http.HandleFunc("/buddystring", exercises.BuddyString)
	// Default handler expects word list for trie
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

func diagonalHandler(w http.ResponseWriter, r *http.Request) {
	var inputArray exercises.TwoDArray
	err := json.NewDecoder(r.Body).Decode(&inputArray)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&inputArray)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output := exercises.Diagonal(inputArray)
	res, _ := json.Marshal(output)
	w.Write(res)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Header().Add("Access-Control-Max-Age", "86400")
	if r.Method == "POST" {
		getMatchingWordList(w, r)
	} else if r.Method == "OPTIONS" {
		res, _ := json.Marshal("Allow: Options, Post")
		w.Write(res)
	}
}

func getMatchingWordList(w http.ResponseWriter, r *http.Request) {
	var wordInfo WordInfo
	err := json.NewDecoder(r.Body).Decode(&wordInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := buildTrie(wordInfo.Words)
	answer := t.GetWordsThatStartWith(wordInfo.Prefix)
	json.NewEncoder(w).Encode(ReturnValues{
		Errcode: 0,
		Message: "OK",
		Data:    answer,
	})
}

func buildTrie(wordList []string) trie.Trie {
	t := trie.TrieData()
	for _, word := range wordList {
		t.Insert(word)
	}
	return *t
}
