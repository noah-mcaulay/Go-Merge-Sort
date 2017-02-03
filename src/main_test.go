package main

import "testing"
import "math/rand"
import "sort"


// TODO: Use table driven test set to make this function "elegant"
// TODO: Split unit tests into own folder and configure Gogland

// Tests for MergeSort function
// (tests divide and merge as well, so for time purposes I will just use this single func)
func TestMergeSort(t *testing.T){

	rand.Seed(42)

	unsorted1K  := [1000]int{}
	sorted1K    := []int{}

	unsorted3   := [3]int{7, 5, 3}
	sorted3     := []int{}

	unsorted2   := [2]int{3, 1}
	sorted2     := []int{}

	unsorted1   := [1]int{8}
	sorted1     := []int{}

	unsorted0   := []int{}
	sorted0     := []int{}

	// Populate unsorted1K with random values
	for i := 0; i < len(unsorted1K); i++ {
		unsorted1K[i] = rand.Int()
	}

	// Get sorted arrays
	sorted1K = MergeSort(unsorted1K[0:])
	sorted3  = MergeSort(unsorted3[0:])
	sorted2  = MergeSort(unsorted2[0:])
	sorted1  = MergeSort(unsorted1[0:])
	sorted0  = MergeSort(unsorted0[0:])

	// Test cases

	// Testing sorted1K
	if !sort.IntsAreSorted(sorted1K) {
		t.Error("Failed: Array of 1000 random ints was improperly sorted.")
	}

	if len(unsorted1K) != len(sorted1K) {
		t.Error("Failed: Array of 1000 random ints was returned with a different length.")
	}


	// Testing sorted3
	if !sort.IntsAreSorted(sorted3) {
		t.Error("Failed: Array of 3 ints was improperly sorted.")
	}

	if len(unsorted3) != len(sorted3) {
		t.Error("Failed: Array of 3 ints was returned with a different length.")
	}


	// Testing sorted2
	if !sort.IntsAreSorted(sorted2) {
		t.Error("Failed: Array of 2 ints was improperly sorted.")
	}

	if len(unsorted2) != len(sorted2) {
		t.Error("Failed: Array of 2 ints was returned with a different length.")
	}


	// Testing sorted3
	if !sort.IntsAreSorted(sorted2) {
		t.Error("Failed: Array of 2 ints was improperly sorted.")
	}

	if len(unsorted2) != len(sorted2) {
		t.Error("Failed: Array of 2 ints was returned with a different length.")
	}


	// Testing sorted1
	if unsorted1[0] != sorted1[0] {
		t.Error("Failed: Array of 1 int was not returned as the same int.")
	}

	if len(unsorted1) != len(sorted1) {
		t.Error("Failed: Array of 1 int was returned with a different length.")
	}


	// Testing sorted0
	if len(sorted0) != 0 {
		t.Error("Failed: Array of 0 ints was returned non-empty.")
	}

}