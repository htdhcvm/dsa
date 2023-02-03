package duplicate_zeros

func DuplicateZeros(arr []int) {
  var lenArr = len(arr);
	var tmp = make([]int, lenArr)

	var j = 0


	for i := 0; i < lenArr; i++ {
		if j == lenArr {
			break
		}
		var current = arr[i]

		if current == 0 {
			tmp[j] = current;
			if j+ 1 <lenArr {
				tmp[j + 1] = current
				j +=1
			}
		} else {
			tmp[j] = current
		}

		j +=1
	}


	for i := 0; i < lenArr; i++ {
		arr[i] = tmp[i]
	}
}

func DuplicateZerosV2(arr []int) {
	var possibleDubs = 0
	var lenArr = len(arr)

	for i := 0; i < lenArr - possibleDubs; i++ {
		var current = arr[i]

		if current == 0 {
			if i == lenArr - possibleDubs - 1 {
				arr[lenArr] = 0;
				lenArr--
				break;
			}
			possibleDubs++;
		}
	}

	var lenArray = lenArr - 1- possibleDubs
	var lenA = len(arr) - 1

	for i := lenArray; i >= 0; i-- {
		var current = arr[i]

		if current == 0 && possibleDubs > 0{
			possibleDubs--
			arr[lenA] = current
			arr[lenA - 1] = current
			lenA--;
		} else {
			arr[lenA] = current
		}

		lenA--;
	}
}