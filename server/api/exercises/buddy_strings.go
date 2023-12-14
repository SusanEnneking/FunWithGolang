package exercises

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Buddies struct {
	S    string
	Goal string
}

func BuddyString(w http.ResponseWriter, r *http.Request) {
	var buddies Buddies
	fmt.Println("body", r.Body)
	err := json.NewDecoder(r.Body).Decode(&buddies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(buddyCheck(buddies.S, buddies.Goal))
}

func buddyCheck(s string, goal string) bool {
	lenS := len(s)
	lenGoal := len(goal)
	if lenS != lenGoal || lenS == 1 {
		return false
	}
	fmt.Println("Length s: ", lenS, " Length goal: ", lenGoal)
	if lenS != lenGoal {
		return false
	}
	// if any two letters are the same, they could be switched to come up with the same string
	// otherwise no
	if s == goal {
		dups := []byte{}
		for index := range s {
			alreadyThere := false
			for dupsIndex := range dups {
				if dups[dupsIndex] == s[index] {
					alreadyThere = true
				}
			}
			if !alreadyThere {
				dups = append(dups, s[index])
			} else {
				return true
			}

		}
		return false
	}

	first, second := -1, -1
	for index := range s {
		if s[index] != goal[index] {
			if first == -1 {
				first = index
			} else if second == -1 {
				second = index
			} else {
				// this is the third mismatch, so the strings cannot be buddies
				return false
			}
		}
	}
	if first == -1 || second == -1 {
		return false
	}

	fmt.Println(first, second)
	fmt.Println(s[first], goal[first], s[second], goal[second])
	return s[first] == goal[second] && s[second] == goal[first]
}
