package squares_sorted_array

import "testing"
import "fmt"

func TestSortedSquares(t *testing.T) {
	var nums = []int{-4,-1,0,3,10}
	var max = SortedSquares(nums)

	fmt.Print(max);
}

