package numbers_even_number

import "testing"

func TestFindNumbers(t *testing.T) {
	var expect = 2
	var nums = []int{12,345,2,6,7896}
	var max = FindNumbers(nums)

	if max != expect { 
		t.Errorf("Expect %d but got %d", expect, max)
	}
}

func Test_2FindNumbers(t *testing.T) {
	var expect = 1
	var nums = []int{555,901,482,1771}
	var max = FindNumbers(nums)

	if max != expect { 
		t.Errorf("Expect %d but got %d", expect, max)
	}
}
