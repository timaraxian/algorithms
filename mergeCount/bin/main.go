package main

import (
	"os"
	"bufio"
	"strconv"

	"github.com/timaraxian/algorithms/mergeCount"
	"fmt"
)

func main() {
	file, err := os.Open("../arraylist.txt")
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

	sorted, inversions := mergeCount.MergeCount(numbers, 0)

	fmt.Println(sorted, "inv:", inversions)
}
