package main

import (
"os"
"bufio"
"strconv"
"fmt"
"github.com/timaraxian/algorithms/quicksort"
)

func main() {
	file, err := os.Open("../quicksort.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	numbers := make([]int, 0)

	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		number, err := strconv.Atoi(string(bytes))
		if err != nil {
			panic(number)
		}

		numbers = append(numbers, number)
	}

	_, comparisons := quicksort.QuickSort(numbers, len(numbers), 0)

	fmt.Println("comp:", comparisons)
}
