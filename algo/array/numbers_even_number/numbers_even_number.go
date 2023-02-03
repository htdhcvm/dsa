package numbers_even_number

import "strconv"

func FindNumbers(nums []int) int {
    var count = 0

		for i := 0; i < len(nums); i++ {
			var item = nums[i];
			var itemString = strconv.Itoa(item)

			if len(itemString) % 2 == 0 {
				count+=1
			}
		}

		return count
}