
package squares_sorted_array



func SortedSquares(nums []int)[]int{
	var tmp = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		var item = nums[i];

		if nums[i] < 0 {
			item = item * -1;
		}

		tmp[i] = item * item;
	}

	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp) - 1; j++ {
			if tmp[j] > tmp [j + 1] {
				var t = tmp[j];
				tmp[j] = tmp[j + 1];
				tmp[j + 1] = t
			}
		}
	}

	return tmp
}