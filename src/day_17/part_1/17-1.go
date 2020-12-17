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

func countActiveNeighbors(nodes map[string]rune, x int, y int, z int) int {
	count := 0
	for ix := x - 1; ix <= x+1; ix++ {
		for iy := y - 1; iy <= y+1; iy++ {
			for iz := z - 1; iz <= z+1; iz++ {
				if ix == x && iy == y && iz == z {
					continue
				}
				nodestr := fmt.Sprintf("%d,%d,%d", ix, iy, iz)
				if node, exists := nodes[nodestr]; exists && node == '#' {
					count++
				}
			}
		}
	}

	return count
}

func simulateCycle(nodes map[string]rune, dims *Dimensions) map[string]rune {
	newNodes := make(map[string]rune, 0)
	for x := dims.minX - 1; x <= dims.maxX+1; x++ {
		for y := dims.minY - 1; y <= dims.maxY+1; y++ {
			for z := dims.minZ - 1; z <= dims.maxZ+1; z++ {
				activeNeighbors := countActiveNeighbors(nodes, x, y, z)
				nodestr := fmt.Sprintf("%d,%d,%d", x, y, z)
				if node, exists := nodes[nodestr]; exists && node == '#' {
					// active
					if activeNeighbors != 2 && activeNeighbors != 3 {
						addNode(newNodes, dims, x, y, z, '.')
					} else {
						addNode(newNodes, dims, x, y, z, '#')
					}
				} else {
					// inactive
					if activeNeighbors == 3 {
						addNode(newNodes, dims, x, y, z, '#')
					} else if exists {
						addNode(newNodes, dims, x, y, z, '.')
					}
				}
			}
		}
	}

	return newNodes
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

func addNode(nodes map[string]rune, dims *Dimensions, x int, y int, z int, node rune) {
	nodeCoords := fmt.Sprintf("%d,%d,%d", x, y, z)
	nodes[nodeCoords] = node
	dims.minX = minInt(x, dims.minX)
	dims.minY = minInt(y, dims.minY)
	dims.minZ = minInt(z, dims.minZ)
	dims.maxX = maxInt(x, dims.maxX)
	dims.maxY = maxInt(y, dims.maxY)
	dims.maxZ = maxInt(z, dims.maxZ)
}

func printNodes(nodes map[string]rune, dims Dimensions) {
	for z := dims.minZ; z <= dims.maxZ; z++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("z = %d", z))
		for y := dims.minY; y <= dims.maxY; y++ {
			for x := dims.minX; x <= dims.maxX; x++ {
				node, exists := nodes[fmt.Sprintf("%d,%d,%d", x, y, z)]
				if exists {
					fmt.Print(string(node))
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	dims := Dimensions{0, 0, 0, 0, 0, 0}
	nodes := make(map[string]rune, 0)
	b, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")
	for lineIdx, line := range lines {
		if line == "" {
			continue
		}
		for nodeIdx, node := range line {
			addNode(nodes, &dims, nodeIdx, lineIdx, 0, node)
		}
	}

	for i := 0; i < 6; i++ {
		nodes = simulateCycle(nodes, &dims)
	}

	activeNodes := 0
	for _, v := range nodes {
		if v == '#' {
			activeNodes++
		}
	}

	fmt.Println(activeNodes)
}
