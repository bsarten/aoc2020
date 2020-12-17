package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Dimensions struct {
	minX int
	minY int
	minZ int
	maxX int
	maxY int
	maxZ int
}

func countActiveNeighbors(nodes map[string]struct{}, x int, y int, z int) int {
	count := 0
	for ix := x - 1; ix <= x+1; ix++ {
		for iy := y - 1; iy <= y+1; iy++ {
			for iz := z - 1; iz <= z+1; iz++ {
				if ix == x && iy == y && iz == z {
					continue
				}
				nodestr := fmt.Sprintf("%d,%d,%d", ix, iy, iz)
				if _, exists := nodes[nodestr]; exists {
					count++
				}
			}
		}
	}

	return count
}

func simulateCycle(nodes map[string]struct{}, dims Dimensions) (map[string]struct{}, Dimensions) {
	newNodes := make(map[string]struct{}, 0)
	newDims := Dimensions{0, 0, 0, 0, 0, 0}

	for x := dims.minX - 1; x <= dims.maxX+1; x++ {
		for y := dims.minY - 1; y <= dims.maxY+1; y++ {
			for z := dims.minZ - 1; z <= dims.maxZ+1; z++ {
				activeNeighbors := countActiveNeighbors(nodes, x, y, z)
				nodestr := fmt.Sprintf("%d,%d,%d", x, y, z)
				if _, exists := nodes[nodestr]; exists {
					// active
					if activeNeighbors == 2 || activeNeighbors == 3 {
						addNode(newNodes, &newDims, x, y, z)
					}
				} else {
					// inactive
					if activeNeighbors == 3 {
						addNode(newNodes, &newDims, x, y, z)
					}
				}
			}
		}
	}

	return newNodes, newDims
}

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

func addNode(nodes map[string]struct{}, dims *Dimensions, x int, y int, z int) {
	nodeCoords := fmt.Sprintf("%d,%d,%d", x, y, z)
	nodes[nodeCoords] = struct{}{}
	dims.minX = minInt(x, dims.minX)
	dims.minY = minInt(y, dims.minY)
	dims.minZ = minInt(z, dims.minZ)
	dims.maxX = maxInt(x, dims.maxX)
	dims.maxY = maxInt(y, dims.maxY)
	dims.maxZ = maxInt(z, dims.maxZ)
}

func main() {
	dims := Dimensions{0, 0, 0, 0, 0, 0}
	nodes := make(map[string]struct{}, 0)
	b, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")
	for lineIdx, line := range lines {
		if line == "" {
			continue
		}
		for nodeIdx, node := range line {
			if node == '#' {
				addNode(nodes, &dims, nodeIdx, lineIdx, 0)
			}
		}
	}

	for i := 0; i < 6; i++ {
		nodes, dims = simulateCycle(nodes, dims)
	}

	fmt.Println(len(nodes))
}
