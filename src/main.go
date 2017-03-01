package main

import (
	"fmt"
	"time"
	"math/rand"
	"sort"
)

// Main function that includes a simple print statement for visible verification.
func main() {

	const PROBLEM_SIZE = 10
	const NUM_RUNS = 1

	rand.Seed(42)

	var totalTime time.Duration

	anArray := [PROBLEM_SIZE]int {}

	for run := 0; run < NUM_RUNS; run++ {

		//fmt.Println(cap(anArray), len(anArray))

		for index := 0; index < len(anArray); index++ {
			anArray[index] = rand.Int()
		}

		before := time.Now()

		sortedArray := MergeSort(anArray[0:])

		fmt.Println("Run: ", run, "Duration: ", time.Now().Sub(before))
		totalTime += time.Now().Sub(before)

		if sort.IntsAreSorted(sortedArray[0:]) {
			fmt.Println("Run 0:", run, "is sorted correctly.")
		}
	}

	fmt.Println("The average duration for ", NUM_RUNS, " runs is: ", totalTime / NUM_RUNS)
}

// MergeSort entrance that is passed "unsorted" and returns a sorted slice
func MergeSort(unsorted []int) []int {
	return Divide(unsorted)
}

// Divide splits anArray into slices and then calls Merge
func Divide(anArray []int) []int {

	length := len(anArray)

	middle := length / 2

	if length < 2 {
		return anArray
	}

	left := Divide(anArray[:middle])
	right := Divide(anArray[middle:])

	return Merge(left, right)
}

// Merge sorts the left and right slices and returns the sortedSlice
func Merge(left []int, right []int) []int {

	lenLeft := len(left)
	lenRight := len(right)

	// TODO: Add a global slice that the sorted values are inserted into for increased performance
	sortedSlice := make([]int, lenLeft+lenRight)

	posLeft := 0
	posRight := 0

	for posLeft < lenLeft && posRight < lenRight {

		if left[posLeft] <= right[posRight] {
			sortedSlice[posLeft+posRight] = left[posLeft]
			posLeft++
		} else {
			sortedSlice[posLeft+posRight] = right[posRight]
			posRight++
		}
	}

	for posLeft < lenLeft {
		sortedSlice[posLeft+posRight] = left[posLeft]
		posLeft++
	}

	for posRight < lenRight {
		sortedSlice[posLeft+posRight] = right[posRight]
		posRight++
	}

	return sortedSlice
}
