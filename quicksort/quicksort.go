package quicksort

func QuickSort(inputArr []int, right int, comparisons int) ([]int, int) {

	//comparisons += len(inputArr) - 1
	if len(inputArr) <= 1 {
		comparisons += 0
		return inputArr, comparisons
	}

	if len(inputArr) == 2 {
		comparisons++
		if inputArr[0]<inputArr[1] {
			inputArr[0], inputArr[1] = inputArr[1], inputArr[0]
		}
		return inputArr, comparisons
	}

	choosePivot(inputArr, 2)
	pivot := inputArr[0]
	i := 1

	for j := 1; j < right; j++ {
		if inputArr[j] < pivot {
			inputArr[j], inputArr[i] = inputArr[i], inputArr[j]
			i++
		}
		comparisons += 1
	}

	inputArr[0], inputArr[i-1] = inputArr[i-1], inputArr[0]

	if len(inputArr[0:i-1]) > 0 {
		_, comparisons = QuickSort(inputArr[0:i-1], len(inputArr[0:i-1]), comparisons)
	}

	if len(inputArr[i:right]) > 0 {
		_, comparisons = QuickSort(inputArr[i:right], len(inputArr[i:right]), comparisons)
	}

	return inputArr, comparisons
}

func choosePivot(inputArr []int, choice int) []int {

	//if len(inputArr) <= 2 {
	//	panic(len(inputArr))
	//}

	switch choice {
	case 1: //pic last one
		inputArr[len(inputArr)-1], inputArr[0] = inputArr[0], inputArr[len(inputArr)-1]
	case 2: // pic median of three
		n := len(inputArr) >> 1
		if len(inputArr) & 1 == 0 {
			n -= 1
		}

		// first item
		first := inputArr[0]
		//  middle item
		middle := inputArr[n]
		//  last item
		last := inputArr[len(inputArr)-1]

		//for _, item := range []int{first,}
		//fmt.Println(n,len(inputArr))
		//last is median
		if (last >= first && last <= middle) || (last <= first && last >= middle) {
			inputArr[0], inputArr[len(inputArr)-1] = inputArr[len(inputArr)-1], inputArr[0]
		}
		// middle is median
		if (middle >= first && middle <= last) || (middle <= first && middle >= last) {
			inputArr[0], inputArr[n] = inputArr[n], inputArr[0]
		}
	default:
	}
	return inputArr
}

