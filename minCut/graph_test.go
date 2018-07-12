package mincut


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGraph_Add(t *testing.T) {
  graph := NewGraph()

  graph.Add(1, 2)
  graph.Add(1, 4)
  graph.Add(2, 4)
  graph.Add(2, 3)
  graph.Add(3, 4)

  // adjcList := [][]int{
  //   {1,2,4},
  //   {2,1,4,3},
  //   {3,2,4},
  //   {4,1,2,3},
  // }

  //graph = graphFromList(adjcList)

  ex := map[int]map[int]int{
    1: {
      2: 1,
      4: 1,
      },
    2: {
      1: 1,
      3: 1,
      4: 1,
    },
    3: {
      2: 1,
      4: 1,
    },
    4: {
      1: 1,
      2: 1,
      3: 1,
    },
  }

  if !reflect.DeepEqual(ex, graph.list) {
    t.Errorf("\n%v \n!= \n%v", ex, graph.list)
  }
}

func TestGraph_AdjcCreate(t *testing.T) {
  graph := NewGraph()

  adjcList := [][]int{
    {1,2,4},
    {2,1,4,3},
    {3,2,4},
    {4,1,2,3},
  }

  graph.AdjcCreate(adjcList)

  ex := map[int]map[int]int{
    1: {
      2: 1,
      4: 1,
      },
    2: {
      1: 1,
      3: 1,
      4: 1,
    },
    3: {
      2: 1,
      4: 1,
    },
    4: {
      1: 1,
      2: 1,
      3: 1,
    },
  }

  if !reflect.DeepEqual(ex, graph.list) {
    t.Errorf("\n%v \n!= \n%v", ex, graph.list)
  }
  
}

func TestGraph_Fuse(t *testing.T) {
  graph := NewGraph()

  graph.Add(1, 2)
  graph.Add(1, 4)
  graph.Add(2, 4)
  graph.Add(2, 3)
  graph.Add(3, 4)

  graph.Fuse(1,2)
  ex := map[int]map[int]int{
    1: {
      3: 1,
      4: 2,
    },
    3: {
      1: 1,
      4: 1,
    },
    4: {
      1: 2,
      3: 1,
    },

  }

  if !reflect.DeepEqual(ex, graph.list) {
    t.Errorf("\n%v \n!= \n%v", ex, graph.list)
  }
}

func TestGraph_Fuse2(t *testing.T) {
  graph := NewGraph()

  graph.Add(1, 2)
  graph.Add(1, 4)
  graph.Add(2, 4)
  graph.Add(2, 3)
  graph.Add(3, 4)

  graph.Fuse(1,2)

  graph.Fuse(1,4)

  ex := map[int]map[int]int{
    1: {
      3: 2,
    },
    3: {
      1: 2,
    },
  }

  if !reflect.DeepEqual(ex, graph.list) {
    t.Errorf("\n%v \n!= \n%v", ex, graph.list)
  }

  if cuts, _ := graph.MinCut(); cuts != 2 {
    t.Error(graph.MinCut())
  }
}

func TestGraph_MinCut(t *testing.T) {
  var minCut int
  for i :=0; i<100; i++ {
  
  graph := NewGraph()

  graph.Add(1, 2)
  graph.Add(1, 4)
  graph.Add(2, 4)
  graph.Add(2, 3)
  graph.Add(3, 4)
    for {
      cuts, err := graph.MinCut()

      if err == nil {
        if i == 0 {
          minCut = cuts
        } else if cuts < minCut{
          minCut = cuts
        }
        //fmt.Println(cuts)
        break
      }

      randA, randB := graph.RandEdge()
      graph.Fuse(randA, randB)

    }
  }
}

func TestGraph_MinCut2(t *testing.T) {

  adjcList := [][]int{
    {1,2,4},
    {2,1,4,3},
    {3,2,4},
    {4,1,2,3},
  }
  var minCut int
  for i :=0; i<100; i++ {
  
    graph := NewGraph()

    for _, row := range(adjcList) {
      for _, v := range(row[1:]) {
        graph.Add(row[0],v)
      }
    }

    for {
      cuts, err := graph.MinCut()

      if err == nil {
        if i == 0 {
          minCut = cuts
        } else if cuts < minCut{
          minCut = cuts
        }
       // fmt.Println(cuts)
        break
      }

      randA, randB := graph.RandEdge()
      graph.Fuse(randA, randB)

    }
  }
  //fmt.Println("min:",minCut)
}

func graphFromList(adjcList [][]int)Graph {
  graph := NewGraph()

  for _, row := range(adjcList) {
      for _, v := range(row[1:]) {
        graph.Add(row[0],v)
      }
  }
  return graph
}