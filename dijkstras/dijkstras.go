package dijkstras

import (
	"github.com/hennersz/datastructures/graph"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

type dijkstrasNode struct {
	name      int
	adjacent  []graph.Edge
	visited   bool
	previous  int
	totalDist int
}

type dijkstrasGraph struct {
	nodes []dijkstrasNode
	edges []graph.Edge
}

func createDijkstrasNode(node graph.Node, startNode int) dijkstrasNode {
	var dist int
	if startNode == node.Name {
		dist = 0
	} else {
		dist = maxInt
	}
	return dijkstrasNode{
		node.Name,
		node.Adjacent,
		false,
		-1,
		dist,
	}
}
func enhanceGraph(inputGraph graph.Graph, start int) dijkstrasGraph {
	var nodes []dijkstrasNode
	for _, node := range inputGraph.Nodes {
		nodes = append(nodes, createDijkstrasNode(node, start))
	}
	return dijkstrasGraph{
		nodes,
		inputGraph.Edges,
	}
}

func findNext(in dijkstrasGraph) (int, bool) {
	index := -1
	lowestDist := maxInt

	for i, node := range in.nodes {
		if lowestDist >= node.totalDist && !node.visited {
			index = i
			lowestDist = node.totalDist
		}
	}

	if index != -1 {
		return index, true
	} else {
		return index, false
	}
}

//ShortestDistance finds the shortest distance between nodes start and end in graph input
func ShortestDistance(input graph.Graph, start, end int) int {
	g := enhanceGraph(input, start)
	for {
		currentNodeIndex, nodeFound := findNext(g)
		if !nodeFound {
			break
		}
		currentNode := &g.nodes[currentNodeIndex]
		currentNode.visited = true
		for _, edge := range currentNode.adjacent {
			nextNode := &g.nodes[edge.End]
			if currentNode.totalDist+edge.Weight < nextNode.totalDist {
				nextNode.totalDist = currentNode.totalDist + edge.Weight
				nextNode.previous = currentNode.name
			}
		}
	}

	return g.nodes[end].totalDist
}
