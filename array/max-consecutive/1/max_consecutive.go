package max_consecutive_1


func FindMaxConsecutiveOnes(nums []int) int {
	var max = 0;
	var currentMax = 0;

	for i := 0; i < len(nums); i++ {
		var current = nums[i];

		if current == 1 {
			currentMax++
			if currentMax > max {
				max = currentMax
			}
		} else {
			currentMax = 0
		}
	}

	return max
}
