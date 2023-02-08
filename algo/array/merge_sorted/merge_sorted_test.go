package merge_sorted

import "testing"
import "reflect"

// import "fmt"

func TestDuplicateZeros(t *testing.T) {
	var expect = []int{1,2,2,3,5,6}
	var nums1 = []int{1,2,3,0,0,0}
	var m = 3
	var nums2 = []int{2,5,6}
	var n = 3

	Merge(nums1, m, nums2, n)


	if ! reflect.DeepEqual(nums1, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZerosMergeTwoPoints(t *testing.T) {
	var expect = []int{1,2,2,3,5,6}
	var nums1 = []int{1,2,3,0,0,0}
	var m = 3
	var nums2 = []int{2,5,6}
	var n = 3

	MergeTwoPoints(nums1, m, nums2, n)


	if ! reflect.DeepEqual(nums1, expect) {
		t.Errorf("Arrays not equal")
	}
}
