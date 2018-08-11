package main

import (
"os"
"bufio"
"strings"
"strconv"
"fmt"
"github.com/timaraxian/algorithms/minCut"
)

func main() {
  file, err := os.Open("mincutArray.txt")
  if err != nil {
    panic(err)
  }

  reader := bufio.NewReader(file)
  adjcList := [][]int{}
  for {
    bytes, _, err := reader.ReadLine()
    if err != nil {
      break
    }

    numbers := strings.Fields(string(bytes))
    intNumb := make([]int, len(numbers))

    for i := 0; i < len(numbers); i++ {
      intNumb[i],_ = strconv.Atoi(numbers[i])
    }

    adjcList = append(adjcList, intNumb)
  }

  var minCut int

  graph := mincut.NewGraph()
  graph.AdjcCreate(adjcList)

  for i :=0; i<100000; i++ {
  
  graph.Reset()

    for {
      cuts, err := graph.MinCut()

      if err == nil {
        if i == 0 {
          minCut = cuts
          fmt.Println(minCut)
        } else if cuts < minCut{
          minCut = cuts
          fmt.Println(minCut)
        } else if i%100 == 0 {
          fmt.Print(".")
        } 

        if minCut == 17 {
          fmt.Println("cuts: ", minCut)
          panic(i)
        }
        break
      }

      randA, randB := graph.RandEdge()
      graph.Fuse(randA, randB)

    }
  }

  fmt.Println("cuts: ", minCut)

}
