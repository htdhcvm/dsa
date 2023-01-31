package max_consecutive_1

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var expect = 3
	var nums = []int{1, 1, 0, 1, 1, 1}
	var max = FindMaxConsecutiveOnes(nums)

	if max != expect { 
		t.Errorf("Expect %d but got %d", expect, max)
	}
}
