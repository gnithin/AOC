package main

import (
	"fmt"
)

type NodeState int

const (
	NODE_CLEAN NodeState = iota
	NODE_INFECTED
)

const (
	STR_NODE_CLEAN    = "."
	STR_NODE_INFECTED = "#"
)

type Position struct {
	row int
	col int
}

type Grid struct {
	intialInfectedPositions map[Position]int
	infectedPositions       map[Position]int
}

func createGridFromInitialMap(gridMap [][]string) Grid {
	grid := Grid{
		make(map[Position]int),
		make(map[Position]int),
	}
	grid.getInfectedNodesFromMap(gridMap)
	return grid
}

func (self *Grid) getInfectedNodesFromMap(gridMap [][]string) {
	height := len(gridMap)
	width := len(gridMap[0])

	if height%2 == 0 || width%2 == 0 {
		panic("Cannot find middle element in even number dimensions")
	}

	center := Position{
		row: height / 2,
		col: width / 2,
	}

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			nodeVal := gridMap[r][c]
			if nodeVal == STR_NODE_INFECTED {
				infPos := Position{
					row: center.row - r,
					col: c - center.col,
				}
				self.infectedPositions[infPos] = 1
			}
		}
	}

	// Copying the initial values
	for key, val := range self.infectedPositions {
		self.intialInfectedPositions[key] = val
	}
}

func (self *Grid) InfectNodeAtPos(pos Position) {
	val, _ := self.infectedPositions[pos]
	self.infectedPositions[pos] = val + 1
}

func (self *Grid) CleanNodeAtPos(pos Position) {
	_, found := self.infectedPositions[pos]
	if found {
		delete(self.infectedPositions, pos)
	}
}

func (self *Grid) GetNodeStateForPos(pos Position) NodeState {
	_, keyFound := self.infectedPositions[pos]
	if keyFound == false {
		return NODE_CLEAN
	}
	return NODE_INFECTED
}

func (self Grid) String() string {
	opStr := ""
	for key, _ := range self.infectedPositions {
		opStr += fmt.Sprintln(key)
	}
	return opStr
}
