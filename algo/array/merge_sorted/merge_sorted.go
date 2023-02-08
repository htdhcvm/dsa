package merge_sorted

// import "fmt"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
	}

	var q = 0
	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[q]
		q++
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1) - 1; j++ {
			if nums1[j] > nums1[j+1] {
				var tmp = nums1[j];

				nums1[j] = nums1[j + 1];
				nums1[j + 1] = tmp
			}
		}
	}
}


func MergeTwoPoints(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
	}

	var res = []int{}

	var i = 0
	var j = 0

	for i < m && j < n {
		if nums1[i] < nums2[j]{
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	for i < m {
		res = append(res, nums1[i])
		i++
	}

	for j < n {
		res = append(res, nums2[j])
		j++
	}

	for i := 0; i < len(res); i++ {
		nums1[i] = res[i]
	}

}