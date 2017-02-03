package GoMergeSort

import "fmt"

func main() {
	fmt.Println("Hello World")

	a := [11]int{6, 2, 3, 5, 10, 8, 12, 1, 0, 7, 4}

	b := []int{}

	b = MergeSort(a[0:])

	fmt.Println("a", a)
	fmt.Println("b", b)
}

func MergeSort(unsorted []int) []int {
	return Divide(unsorted)
}

func Divide(anArray []int) []int {

	length := len(anArray)

	middle := length / 2

	if length < 2 {
		return anArray
	}

	left :=  Divide(anArray[:middle])
	right := Divide(anArray[middle:])

	return Merge(left, right)
}

func Merge(left []int, right []int) []int {

	lenLeft   := len(left)
	lenRight  := len(right)

	sortedSlice := make([]int, lenLeft + lenRight)

	posLeft   := 0
	posRight  := 0

	for posLeft < lenLeft && posRight < lenRight {

		if left[posLeft] <= right[posRight] {
			sortedSlice[posLeft + posRight] = left[posLeft]
			posLeft++
		} else {
			sortedSlice[posLeft + posRight] = right[posRight]
			posRight++
		}
	}

	for posLeft < lenLeft {
		sortedSlice[posLeft + posRight] = left[posLeft]
		posLeft++
	}

	for posRight < lenRight {
		sortedSlice[posLeft + posRight] = right[posRight]
		posRight++
	}

	return sortedSlice
}
