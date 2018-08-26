package main

import (
	"github.com/timaraxian/algorithms/heaps"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	minHeap := heaps.NewHeap()
	maxHeap := heaps.NewHeap()
	sumMedian := 0

	fmt.Println("opening file...")
	file, err := os.Open("median.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Reading file...")
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		num, err := strconv.Atoi(strings.Trim(string(line), " "))
		if err != nil {
			panic(num)
		}

		if len(maxHeap.Items) == 0 {
			maxHeap.Insert(-num)
			sumMedian += -maxHeap.Items[0]
			continue
		} else if len(minHeap.Items) == 0 {
			if -maxHeap.Items[0] > num {
				firstNum, ok := maxHeap.ExtractMin()
				if !ok {
					panic(firstNum)
				}
				minHeap.Insert(-firstNum)
				maxHeap.Insert(-num)
			} else {
				minHeap.Insert(num)
			}
			sumMedian += -maxHeap.Items[0]
			continue
		}

		if num < -maxHeap.Items[0] {
			maxHeap.Insert(-num)
		} else {
			minHeap.Insert(num)
		}
		if len(maxHeap.Items)-len(minHeap.Items) >= 2 {
			swap, ok := maxHeap.ExtractMin();
			if !ok {
				panic(swap)
			}
			minHeap.Insert(-swap)
		}
		if len(minHeap.Items)-len(maxHeap.Items) >= 2 {
			swap, ok := minHeap.ExtractMin();
			if !ok {
				panic(swap)
			}
			maxHeap.Insert(-swap)
		}
		if len(minHeap.Items) > len(maxHeap.Items) {
			sumMedian += minHeap.Items[0]
		} else {
			sumMedian += -maxHeap.Items[0]
		}
	}
	fmt.Println("sum of medians mod 10000:",sumMedian % 10000)
}
