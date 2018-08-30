package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

func main() {
	numbs := map[int]bool{}
	numbsList := make([]int, 0)
	fmt.Println("opening file...")
	file, err := os.Open("hash.txt")
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
		if _, ok := numbs[num]; !ok {
			numbs[num] = true
			numbsList = append(numbsList,num)
		}
	}
	fmt.Println("sorting numbs...")
	sort.Ints(numbsList)

	found := 0
	targets := map[int]bool{}
	for i := -10000; i <= 10000; i++ {
		if _, ok := targets[i]; !ok {
			targets[i] = false
		}
	}
	t := len(numbsList)/2
	for j := 0; j < len(numbsList); j++ {
		num := numbsList[j]
		if t == 0 {
			break
		}
		for i := -10000; i <= 10000; i++ {
			if targets[i] {
				continue
			}
			diff := i - num
			if diff == num {
				continue
			}
			if numbs[diff] {
				fmt.Printf("found: %v with %v & %v \n", i, num, diff)
				targets[i] = true
				found ++
				if found % 50 == 0 && found != 0 {
					fmt.Printf("\n\nfound: %v\n\n",found)
				}
			}
		}
		t --
	}
	fmt.Println(found)
}
