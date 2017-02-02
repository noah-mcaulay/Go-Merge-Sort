package main

import "fmt"



func main() {
	fmt.Println("Hello World")

	a := [11]int{6, 2, 3, 5, 10, 8, 12, 1, 0, 7, 4}

	b := []int{}

	b = divide(a[0:])
	//b = a[0:3]

	fmt.Println("a", a)
	fmt.Println("b", b)


}

//sortedArray := []int

func divide(anArray []int) []int {

	length := len(anArray)

	if length < 2 {
		return anArray
	}

	fmt.Println("anArray", anArray)
	left :=  divide(anArray[:length/2])
	right := divide(anArray[length/2:])

	//divide(left)
	//divide(right)

	//fmt.Println("anArray", anArray)
	//fmt.Println("len:", length)
	//fmt.Println("len/2:", length/2)


	//return anArray
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	/*if len(anArray) == 0 {
		return anArray
	}*/
	lenLeft   := len(left)
	lenRight  := len(right)

	posSorted := 0
	posLeft   := 0
	posRight  := 0

	fmt.Println("left", left)
	fmt.Println("right", right)
	//fmt.Println("sortedarray", sortedArray)

	var sortedArray []int

	for posLeft < lenLeft && posRight < lenRight {

		if left[posLeft] > right[posRight] {
			sortedArray[posSorted] = right[posRight]
			posRight++
		} else {
			sortedArray[posSorted] = left[posLeft]
			posRight++
		}
		posSorted++
	}

	for posLeft <= lenLeft {
		sortedArray[posSorted] = right[posRight]
		posRight++
		posSorted++
	}

	for posRight <= lenRight {
		sortedArray[posSorted] = left[posLeft]
		posLeft++
		posSorted++
	}


	return sortedArray
}
