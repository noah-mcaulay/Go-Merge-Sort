package main

import "fmt"



func main() {
	fmt.Println("Hello World")

	a := [11]int{6, 2, 3, 5, 10, 8, 12, 1, 0, 7, 4}

	//b := [10]int{}

	divide(a[0:])
	fmt.Println(divide(a[0:]))
	fmt.Println(a)


}

//sortedArray := []int

func divide(anArray []int) []int {

	length := len(anArray)

	if length == 0 {
		return anArray
	}

	fmt.Println("len:", length)
	fmt.Println("len/2:", length/2)

	//return anArray
	return merge(anArray[0:length/2], anArray[length/2:])
}

func merge(lArray []int, rArray []int) []int {
	/*if len(anArray) == 0 {
		return anArray
	}*/
	sortedArray := []int{}

	sortedArray

	return lArray
}
