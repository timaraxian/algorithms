package mincut

import (
  "errors"
  "math/rand"
  "time"
)

type Vertex int 

type Graph struct {
  list map[Vertex]map[Vertex]int
  listcopy map[Vertex]map[Vertex]int

}

func NewGraph() Graph{
  return Graph{
    list: make(map[Vertex]map[Vertex]int),
    listcopy: make(map[Vertex]map[Vertex]int),
  } 
}

func (graph *Graph) Add(a, b Vertex) {
  
  graph.add(a,b,1)
}

func (graph *Graph) add(a, b Vertex, incr int) {
  
  if _, ok := graph.list[a]; !ok {
   graph.list[a] = make(map[Vertex]int) 
  }

  if _, ok := graph.list[a][b]; !ok {
    graph.list[a][b] = incr
  } else {
    graph.list[a][b] += incr 
  }

  if _, ok := graph.list[b]; !ok {
   graph.list[b] = make(map[Vertex]int) 
  }

  if _, ok := graph.list[b][a]; !ok {
    graph.list[b][a] = incr
  } else {
    graph.list[b][a] += incr
  }
}

func (graph *Graph) AdjcCreate(adjcList [][]int) {
  // for row := 0; row < len(list); row++ {
  //   for item :=0; item <len(list[row]); item++ {
  //     if item == 0 {
  //       continue
  //     }
  //     graph.Add(Vertex(list[row][0]),Vertex(list[row][item]))
  //   }
  // }

  // for key,_ := range(graph.list) {
  //   for key1,_ := range(graph.list[key]) {
  //     graph.list[key][key1] = graph.list[key][key1]/2
  //   }
  // }
  for _,row := range(adjcList) {
    a := Vertex(row[0])
    graph.list[a] = make(map[Vertex]int)
    for _,bStr := range(row[1:]) {
      b := Vertex(bStr)
      if _,ok := graph.list[a][b]; ok {
        panic(b)
      }
      graph.list[a][b] = 1
    } 
  }
  copyList(graph.list, graph.listcopy)
}

func copyList(listA, listB map[Vertex]map[Vertex]int) {
  for a,bMap := range(listA) {
    listB[a] = make(map[Vertex]int)
    for b := range(bMap) {
      listB[a][b] = 1
    } 
  }
}

func (graph *Graph) Reset() {
  copyList(graph.listcopy, graph.list)
}

func (graph *Graph) Fuse(a, b Vertex) {

  if _, ok := graph.list[a]; !ok {
    panic("Edge doesn't exist")
  }
  if _, ok := graph.list[a][b]; !ok {
    panic("Edge doesn't exist")
  }

  // loop over keys in vertexB
  for key,value := range(graph.list[b]) {
    //check to see if key is a
    if key == a {
      continue
    } 

    graph.add(a, key, value)
  }

  // delete vB from Graph
  delete(graph.list, b)

  for key := range(graph.list) {
    delete(graph.list[key],b)
  }
}

// func (graph Graph) RandVertex()int {
//   stop := rand.Intn(len(graph.list))
//   i := 0
//   for key := range(graph.list) {
//     if i == stop {
//       return key
//     }
//     i++
//   }
//   return 0
// }

func (graph Graph) RandEdge()(Vertex, Vertex){
  rand.Seed(time.Now().UnixNano())
  // fmt.Println(time.Now().UnixNano())
  stopA := rand.Intn(len(graph.list))
  i := 0
  for keyA := range(graph.list) {
    if i == stopA {
      stopB := rand.Intn(len(graph.list[keyA]))
      j := 0
      for keyB := range(graph.list[keyA]) {
        if j == stopB {
          return keyA, keyB
        }
        j++
      }
    }
    i++
  }
  return 0,0
}

func (graph Graph) MinCut() (int, error) {
  if len(graph.list) != 2 {
    return 0, errors.New("Too many nodes")
  }
  for _, b := range(graph.list) {
    for _, cuts := range(b) {
      return cuts, nil
    }
  }
  return 0, nil
}