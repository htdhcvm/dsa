package merge_sort

import "testing"
import "reflect"

func TestMergeSort(t *testing.T) {
	var example = []int{12, 11, 13, 5, 6, 7}
	var expected = []int{5, 6, 7, 11, 12, 13}
	
	var sorted = MergeSort(example)

	if ! reflect.DeepEqual(expected, sorted) {
		t.Errorf("Arrays not equal")
	}
}