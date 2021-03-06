package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func minInt(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func maxInt(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

type Dimensions struct {
	min Coordinates
	max Coordinates
}

type Coordinates struct {
	x int
	y int
	z int
}

type Empty struct {
}

type Nodes struct {
	dims      Dimensions
	dimsValid bool
	nodes     map[Coordinates]Empty
}

func (n *Nodes) exists(coords Coordinates) bool {
	_, exists := n.nodes[coords]
	return exists
}

func newNodes() *Nodes {
	n := Nodes{nodes: make(map[Coordinates]Empty, 0)}
	return &n
}

func (n *Nodes) addNode(coords Coordinates) {
	n.nodes[coords] = Empty{}
	if n.dimsValid {
		n.dims.min.x = minInt(coords.x, n.dims.min.x)
		n.dims.min.y = minInt(coords.y, n.dims.min.y)
		n.dims.min.z = minInt(coords.z, n.dims.min.z)
		n.dims.max.x = maxInt(coords.x, n.dims.max.x)
		n.dims.max.y = maxInt(coords.y, n.dims.max.y)
		n.dims.max.z = maxInt(coords.z, n.dims.max.z)
	} else {
		n.dims.min = coords
		n.dims.max = coords
		n.dimsValid = true
	}
}

func (n *Nodes) countActiveNeighborsAt(coords Coordinates) int {
	count := 0
	var checkCoords Coordinates
	for checkCoords.x = coords.x - 1; checkCoords.x <= coords.x+1; checkCoords.x++ {
		for checkCoords.y = coords.y - 1; checkCoords.y <= coords.y+1; checkCoords.y++ {
			for checkCoords.z = coords.z - 1; checkCoords.z <= coords.z+1; checkCoords.z++ {
				if checkCoords.x == coords.x && checkCoords.y == coords.y && checkCoords.z == coords.z {
					continue
				}
				if n.exists(checkCoords) {
					count++
				}
			}
		}
	}

	return count
}

func simulateCycle(nodes *Nodes) *Nodes {
	newNodes := newNodes()

	var coords Coordinates
	for coords.x = nodes.dims.min.x - 1; coords.x <= nodes.dims.max.x+1; coords.x++ {
		for coords.y = nodes.dims.min.y - 1; coords.y <= nodes.dims.max.y+1; coords.y++ {
			for coords.z = nodes.dims.min.z - 1; coords.z <= nodes.dims.max.z+1; coords.z++ {
				activeNeighbors := nodes.countActiveNeighborsAt(coords)
				if nodes.exists(coords) {
					// active
					if activeNeighbors == 2 || activeNeighbors == 3 {
						newNodes.addNode(coords)
					}
				} else {
					// inactive
					if activeNeighbors == 3 {
						newNodes.addNode(coords)
					}
				}
			}
		}
	}

	return newNodes
}

func printNodes(nodes *Nodes) {
	for z := nodes.dims.min.z; z <= nodes.dims.max.z; z++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("z = %d", z))
		for y := nodes.dims.min.y; y <= nodes.dims.max.y; y++ {
			for x := nodes.dims.min.x; x <= nodes.dims.max.x; x++ {
				if nodes.exists(Coordinates{x, y, z}) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	const numCycles = 6
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nodes := newNodes()
	lines := strings.Split(string(b), "\n")
	for lineIdx, line := range lines {
		if line == "" {
			continue
		}
		for nodeIdx, node := range line {
			if node == '#' {
				nodes.addNode(Coordinates{nodeIdx, lineIdx, 0})
			}
		}
	}

	for i := 0; i < numCycles; i++ {
		nodes = simulateCycle(nodes)
	}

	fmt.Println(len(nodes.nodes))
}
