package mergeCount

// start with an array A of length n
func MergeCount(input []int, invCount int) (D []int, inversions int) {
	// base case n == 1
	D = make([]int, len(input))
	inversions = invCount
	if len(input) == 1 {
		D = input
		return D, inversions
	}

	// split the array into sub arrays B and C  at index n/2
	B, C := Splitter(input, len(input))

	// call the sort and count function on each sub array (recursive)
	if len(B) > 2 {
		B, inversions = MergeCount(B, inversions)
	}
	if len(C) > 2 {
		C, inversions = MergeCount(C, inversions)
	}
	if len(B) == 2 && (B[0]>B[1]) {
		B = Switch(B)
		inversions++
	}
	if len(C) == 2 && (C[0]>C[1]) {
		C = Switch(C)
		inversions++
	}

	//merge the two arrays together, adding to the count whenever there is an add from C whilst B still exists
	i := 0
	j := 0
	for k := 0; k < len(D); k++ {
		if !(i == len(B)) && (j == len(C) || B[i] < C[j]) {
			D[k] = B[i]
			if j >= i {
				inversions+=0
			}
			i++
		} else {
			D[k] = C[j]
			j++
			if !(i == len(B)) {
				inversions += len(B) - i
			}
		}
	}
	return D, inversions
}

func Splitter(input []int, n int) ([]int, []int) {
	n = n >> 1
	B := input[:n]
	C := input [n:]
	return B, C
}

func Switch(input []int) ([]int) {
	input[0], input[1] = input[1], input[0]
	return input
}
