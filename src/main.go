package main

import (
	"fmt"
	"time"
	"math/rand"
	"sort"
)

// Main function that includes simple benchmarking.
func main() {

	const PROBLEM_SIZE = 10000000
	const NUM_RUNS = 10

	rand.Seed(42)

	var totalTime time.Duration

	anArray := [PROBLEM_SIZE]int {}

	for run := 0; run < NUM_RUNS; run++ {

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

	fmt.Println("The average duration for ", NUM_RUNS, " runs of ", PROBLEM_SIZE, " is: ", totalTime / NUM_RUNS)
}

// MergeSort entrance that is passed "unsorted" and returns a sorted slice (also checks if the passed in array is too small)
func MergeSort(unsorted []int) []int {

	if len(unsorted) < 2 {
		return unsorted
	} else {
		return Divide(unsorted)
	}
}

// Divide splits anArray into slices and then calls Merge
func Divide(anArray []int) []int {

	length := len(anArray)

	middle := length / 2

	if length < 2 {
		return anArray
	}

	// Recursively call Divide on left side of slice
	left := Divide(anArray[:middle])

	// Recursively call Divide on right side of slice
	right := Divide(anArray[middle:])

	// Return the sorted slice of left with right
	return Merge(left, right)
}

// Merge sorts the left and right slices and returns the sortedSlice
func Merge(left []int, right []int) []int {

	lenLeft := len(left)
	lenRight := len(right)

	// sortedSlice is where the sorted values are inserted and then return
	//	(for increased performance there should be a global slice that
	//	 is inserted into to prevent unneeded allocations)
	sortedSlice := make([]int, lenLeft+lenRight)

	// These keep track of the positions as we merge the slices into a sorted slice
	posLeft := 0
	posRight := 0

	// While there are still items to be sorted on the left AND right:
	// 	keep adding them to our sorted slice
	for posLeft < lenLeft && posRight < lenRight {

		if left[posLeft] <= right[posRight] {
			sortedSlice[posLeft+posRight] = left[posLeft]
			posLeft++
		} else {
			sortedSlice[posLeft+posRight] = right[posRight]
			posRight++
		}
	}

	// While there are unsorted values in the left slice, add them to the sorted slice
	for posLeft < lenLeft {
		sortedSlice[posLeft+posRight] = left[posLeft]
		posLeft++
	}

	// While there are unsorted values in the right slice, add them to the sorted slice
	for posRight < lenRight {
		sortedSlice[posLeft+posRight] = right[posRight]
		posRight++
	}

	return sortedSlice
}
