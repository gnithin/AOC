package main

import (
	"fmt"
)

type NodeState int

const (
	NODE_CLEAN NodeState = iota
	NODE_INFECTED
	NODE_WEAKENED
	NODE_FLAGGED
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
	flaggedPositions        map[Position]int
	weakenedPositions       map[Position]int
}

func createGridFromInitialMap(gridMap [][]string) Grid {
	grid := Grid{
		make(map[Position]int),
		make(map[Position]int),
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
	_, found := self.weakenedPositions[pos]
	if !found {
		panic("should not happen!")
	}
	delete(self.weakenedPositions, pos)
	self.infectedPositions[pos] = 1
}

func (self *Grid) WeakenNodeAtPos(pos Position) {
	self.weakenedPositions[pos] = 1
}

func (self *Grid) FlagNodeAtPos(pos Position) {
	_, found := self.infectedPositions[pos]
	if !found {
		panic("should not happen!")
	}
	delete(self.infectedPositions, pos)
	self.flaggedPositions[pos] = 1
}

func (self *Grid) CleanNodeAtPos(pos Position) {
	delete(self.flaggedPositions, pos)
}

func (self *Grid) GetNodeStateForPos(pos Position) NodeState {
	_, isFlagged := self.flaggedPositions[pos]
	if isFlagged {
		return NODE_FLAGGED
	}
	_, isWeakened := self.weakenedPositions[pos]
	if isWeakened {
		return NODE_WEAKENED
	}
	_, isInfected := self.infectedPositions[pos]
	if isInfected {
		return NODE_INFECTED
	}
	return NODE_CLEAN
}

func (self Grid) String() string {
	opStr := ""
	for key, _ := range self.infectedPositions {
		opStr += fmt.Sprintln(key)
	}
	return opStr
}
