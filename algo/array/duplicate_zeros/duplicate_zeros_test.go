package duplicate_zeros

import "testing"
import "reflect"

func TestDuplicateZeros(t *testing.T) {
	var expect = []int{1,0,0,2,3,0,0,4}
	var nums = []int{1,0,2,3,0,4,5,0}

	DuplicateZeros(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZerosAll(t *testing.T) {
	var expect = []int{0,0,0,0,0,0,0}
	var nums = []int{0,0,0,0,0,0,0}

	DuplicateZeros(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZeros_3(t *testing.T) {
	var expect = []int{1,0,0,2,3,0,0,4}
	var nums = []int{1,0,2,3,0,4,5,0}

	DuplicateZeros(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZerosV2(t *testing.T) {
	var expect = []int{1,0,0,2,3,0,0,4}
	var nums = []int{1,0,2,3,0,4,5,0}

	DuplicateZerosV2(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZerosAV2(t *testing.T) {
	var expect = []int{8,4,5,0,0,0,0,0}
	var nums = []int{8,4,5,0,0,0,0,7}

	DuplicateZerosV2(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZerosAllV2(t *testing.T) {
	var expect = []int{0,0,0,0,0,0,0}
	var nums = []int{0,0,0,0,0,0,0}

	DuplicateZerosV2(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZeros_3V2(t *testing.T) {
	var expect = []int{1,0,0,2,3,0,0,4}
	var nums = []int{1,0,2,3,0,4,5,0}

	DuplicateZerosV2(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}

func TestDuplicateZeros_4V2(t *testing.T) {
	var expect = []int{9,9,9,4,8,0,0,0,0,3,7,2,0,0,0,0,0,0,0,0,9,1,0,0,0,0,1,1,0,0,5,6,3,1,6,0,0,0,0,2,3,4,7,0,0,3,9,3,6,5,8,9,1,1,3,2,0,0,0,0,7,3,3,0,0,5,7,0,0,8,1,9,6,3,0,0,8,8,8,8,0}
	var nums = []int{9,9,9,4,8,0,0,3,7,2,0,0,0,0,9,1,0,0,1,1,0,5,6,3,1,6,0,0,2,3,4,7,0,3,9,3,6,5,8,9,1,1,3,2,0,0,7,3,3,0,5,7,0,8,1,9,6,3,0,8,8,8,8,0,0,5,0,0,0,3,7,7,7,7,5,1,0,0,8,0,0}

	DuplicateZerosV2(nums)

	if ! reflect.DeepEqual(nums, expect) {
		t.Errorf("Arrays not equal")
	}
}