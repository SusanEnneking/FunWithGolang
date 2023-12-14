package exercises

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiagonal(t *testing.T) {
	expected := []int{1, 4, 2, 7, 5, 3, 8, 6, 9}
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	twoDArray := TwoDArray{
		Td: nums,
	}
	answer := Diagonal(twoDArray)
	if !reflect.DeepEqual(answer, expected) {
		t.Error(fmt.Printf("Expected: %v, but got: %v.", answer, expected))
	}
	nums = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7},
		{8},
		{9, 10, 11},
		{12, 13, 14, 15, 16},
	}
	expected = []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}
	twoDArray = TwoDArray{
		Td: nums,
	}
	answer = Diagonal(twoDArray)
	if !reflect.DeepEqual(answer, expected) {
		t.Error(fmt.Printf("Expected: %v, but got: %v.", answer, expected))
	}
}
