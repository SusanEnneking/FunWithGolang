package twosum

import (
	"fmt"

	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(TwoSum(input, target))
	//expect [1, 0]

	input = []int{3, 2, 4}
	target = 6
	fmt.Println(TwoSum(input, target))
	//expect [1,2]

	input = []int{3, 3}
	target = 6
	fmt.Println(TwoSum(input, target))
	//expect [0,1]

	input = []int{3}
	target = 3
	fmt.Println(TwoSum(input, target))
	//expect[]

	input = []int{3, 2, 3}
	target = 6
	//expect [0,2]
	fmt.Println(TwoSum(input, target))

	input = []int{-1, -2, -3, -4, -5}
	target = -8
	//expect [2,4]
	fmt.Println(TwoSum(input, target))

	input = []int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1}
	target = 11
	//expect [5,11]
	fmt.Println(TwoSum(input, target))
}
