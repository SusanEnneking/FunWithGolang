package exercises

import (
	"fmt"
)

type Buddies struct {
	S    string
	Goal string
}

func BuddyCheck(buddies Buddies) bool {
	lenS := len(buddies.S)
	lenGoal := len(buddies.Goal)
	if lenS != lenGoal || lenS == 1 {
		return false
	}
	fmt.Println("Length s: ", lenS, " Length goal: ", lenGoal)
	if lenS != lenGoal {
		return false
	}
	// if any two letters are the same, they could be switched to come up with the same string
	// otherwise no
	if buddies.S == buddies.Goal {
		dups := []byte{}
		for index := range buddies.S {
			alreadyThere := false
			for dupsIndex := range dups {
				if dups[dupsIndex] == buddies.S[index] {
					alreadyThere = true
				}
			}
			if !alreadyThere {
				dups = append(dups, buddies.S[index])
			} else {
				return true
			}

		}
		return false
	}

	first, second := -1, -1
	for index := range buddies.S {
		if buddies.S[index] != buddies.Goal[index] {
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
	fmt.Println(buddies.S[first], buddies.Goal[first], buddies.S[second], buddies.Goal[second])
	return buddies.S[first] == buddies.Goal[second] && buddies.S[second] == buddies.Goal[first]
}
