package elementchecker

func ElementChecker(input []int, element int) bool {
	for i = 0; i < len(input); i++ {
		if input[i] == element {
			return true
		}
	}
	return false
}
