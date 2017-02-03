package GoMergeSort

import "testing"
import "math/rand"

func TestMergeSort(t *testing.T){

	rand.Seed(42)

	unsorted := []int{}
	sorted := []int{}

	for i := 0; i < 1000; i++ {
		unsorted[i] = rand.Int()
	}

	sorted = MergeSort(unsorted)

	for i := 0; i < len(sorted); i++ {

		if sorted[i] > sorted[i + 1] {
			t.Error("Incorrectly sorted.")
		}
	}


}