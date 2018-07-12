package algorithms

func MergeSort(input []int) ([]int) {
	output := make([]int, len(input))

	//base case
	if len(input) == 0 {
		return output
	}

	//case: input == an array containing 1 number.
	if len(input) == 1 {
		copy(output, input)
		return output
	}

	// case: input == an array containing 2 numbers
	//if len(input) == 2 {
	//	if input[0] > input[1] {
	//		output[0] = input[1]
	//		output[1] = input[0]
	//	} else {
	//		copy(output, input)
	//	}
	//	return output
	//}

	// case: input is larger than 2 so has to split the array to have recursion
	mid := len(input) >> 1

	left := MergeSort(input[mid:])
	right := MergeSort(input[:mid])

	i := 0
	j := 0
	for k := 0; k < len(output); k++ {
		if !(i == len(left)) && (j == len(right) || left[i] < right[j]) {
			output[k] = left[i]
			i++
		} else {
			output[k] = right[j]
			j++
		}
	}
	return output
}
