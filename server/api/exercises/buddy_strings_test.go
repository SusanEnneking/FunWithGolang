package exercises

import (
	"fmt"
	"testing"
)

func TestBuddyStrings(t *testing.T) {
	expected := true

	buddies := Buddies{
		S:    "ab",
		Goal: "ba",
	}
	answer := BuddyCheck(buddies)
	if answer != expected {
		t.Error(fmt.Printf("Expected: %t, but got: %t, for input %v", answer, expected, buddies))
	}
	expected = false

	buddies = Buddies{
		S:    "ab",
		Goal: "ab",
	}
	answer = BuddyCheck(buddies)
	if answer != expected {
		t.Error(fmt.Printf("Expected: %t, but got: %t, for input %v", answer, expected, buddies))
	}
	expected = true

	buddies = Buddies{
		S:    "aa",
		Goal: "aa",
	}
	answer = BuddyCheck(buddies)
	if answer != expected {
		t.Error(fmt.Printf("Expected: %t, but got: %t, for input %v", answer, expected, buddies))
	}
}
