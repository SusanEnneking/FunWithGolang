package exercises

import (
	"sync"
)

// this is a horrible solution to the problem, but it allows me to use go routines!
// the question goes something like this:  Given a 2-d array return a 1d array with containing all of the diagonals in order
// see test for expected input/output combos

type TwoDArray struct {
	Td [][]int
}

var wg = &sync.WaitGroup{} //pointer to wait group
func Diagonal(inputArray TwoDArray) []int {

	c1 := make(chan []int, 1)
	c2 := make(chan []int, 2)
	wg.Add(1)
	go getDiagonalsStartingInCol0(inputArray.Td, c1)
	longestArrayLength := 0
	for i1 := range inputArray.Td {
		if len(inputArray.Td[i1]) > longestArrayLength {
			longestArrayLength = len(inputArray.Td[i1])
		}
	}
	wg.Add(1)
	go getDiagonalStartingInLastArray(inputArray.Td, longestArrayLength, c2)
	wg.Wait()
	output := <-c1
	output = append(output, <-c2...)
	return output
}

func getDiagonalsStartingInCol0(nums [][]int, out chan<- []int) {
	var output []int
	for i1 := range nums {
		adder := 0
		for index := i1; index >= 0; index-- {
			if adder < len(nums[index]) {
				output = append(output, nums[index][adder])
			}
			adder++
		}
	}

	out <- output
	wg.Done()
}

func getDiagonalStartingInLastArray(nums [][]int, longestArrayLength int, out chan<- []int) {
	lastArrayIndex := len(nums) - 1
	var output []int
	for rowIndex := 1; rowIndex < longestArrayLength; rowIndex++ {
		adder := rowIndex
		for index := lastArrayIndex; index >= 0; index-- {
			if adder < len(nums[index]) {
				output = append(output, nums[index][adder])
			}
			adder++
		}
	}
	out <- output
	wg.Done()
}
