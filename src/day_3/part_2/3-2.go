package main

import (
	"bufio"
	"fmt"
	"os"
)

func findTrees(treeMap []string, right int, down int) int {
	posX := right
	posY := down
	height := len(treeMap)
	var treesEncountered int
	for posY < height {
		if treeMap[posY][posX] == '#' {
			treesEncountered++
		}

		posY += down
		posX = (posX + right) % len(treeMap[0])
	}

	return treesEncountered
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var treeMap []string
	for scanner.Scan() {
		treeMap = append(treeMap, scanner.Text())
	}

	treesEncountered := findTrees(treeMap, 1, 1)
	treesEncountered *= findTrees(treeMap, 3, 1)
	treesEncountered *= findTrees(treeMap, 5, 1)
	treesEncountered *= findTrees(treeMap, 7, 1)
	treesEncountered *= findTrees(treeMap, 1, 2)

	fmt.Println(treesEncountered)
	q := new(Q)
}
