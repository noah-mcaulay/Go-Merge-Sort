package main

import "fmt"

// Main function that includes a simple print statement for visible verification.
func main() {

	a := [11]int{6, 2, 3, 5, 10, 8, 12, 1, 0, 7, 4}

	b := []int{}

	b = MergeSort(a[0:])

	fmt.Println("a", a)
	fmt.Println("b", b)
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
// TODO: See if creating a global slice that the sorted values are placed into will improve performance
func Merge(left []int, right []int) []int {

	lenLeft := len(left)
	lenRight := len(right)

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
