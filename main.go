package main

import (
	"fmt"

	"github.com/hennersz/algorithms/dijkstras"
	"github.com/hennersz/datastructures/graph"
)

func main() {
	data := [][]int{
		{-1, 1, -1},
		{1, -1, 1},
		{-1, 1, -1},
	}

	g := graph.CreateGraph(data)
	res := dijkstras.ShortestDistance(g, 0, 2)
	fmt.Println(res)
}
